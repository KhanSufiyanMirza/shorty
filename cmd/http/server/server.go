package server

import (
	"context"
	"encoding/json"
	"fmt"
	"hex/config"
	fasthttproutes "hex/internal/adapters/inbound/fasthttpl/routes"
	"hex/internal/adapters/inbound/gin/routes"

	"hex/internal/ports/inbound"
	"hex/internal/ports/outbound"
	"hex/utils/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

type Server struct {
	httpServer httpServer
	appConfig  *config.Config
	logger     logger.Logger
	api        inbound.APIPort
	db         outbound.DbPort
}

const RegulatorServerModule = "RegulatorServer"

type httpServer interface {
	Start()
	GracefulShutdown()
}

// TODO: think of a better name
type ginServer struct {
	httpServer *http.Server
	appConfig  *config.Config
	logger     logger.Logger
	api        inbound.APIPort
	db         outbound.DbPort
}
type fasthttpServer struct {
	httpServer *fasthttp.Server
	appConfig  *config.Config
	logger     logger.Logger
	api        inbound.APIPort
}

func New(configs *config.Config, api inbound.APIPort, logger logger.Logger, db outbound.DbPort) *Server {
	var httpServer httpServer

	switch configs.HttpRouter {
	case "fasthttp":
		httpServer = newFastHttpServer(configs, api, logger)
	default:
		httpServer = newGinServer(configs, api, logger, db)
	}

	return &Server{
		httpServer: httpServer,
		appConfig:  configs,
		logger:     logger,
		api:        api,
		db:         db,
	}
}
func (server *Server) Run() {
	// server.seed()
	server.start()
	server.gracefulShutdown()
}

func (server *Server) start() {
	server.httpServer.Start()
}

func (server *Server) gracefulShutdown() {
	server.httpServer.GracefulShutdown()
}

// TODO: seed is used to feed prerequisite data
func (server *Server) seed() {
}

// gin or http server
func newGinServer(configs *config.Config, api inbound.APIPort, logger logger.Logger, db outbound.DbPort) *ginServer {

	return &ginServer{
		appConfig: configs,
		logger:    logger,
		api:       api,
		db:        db,
	}
}

func (server *ginServer) Start() {
	port := fmt.Sprintf(":%v", server.appConfig.AppPort)
	server.httpServer = &http.Server{
		Addr: port,
		// using gin router
		Handler: routes.NewRouter(*server.appConfig, server.api, server.logger).SetRouters(),
	}

	if a := recover(); a != nil {
		var bytes []byte
		bytes, _ = json.Marshal(a)
		server.logger.Error("Server", fmt.Sprintf("recovering from panic %v", string(bytes)), fmt.Errorf("recovering from panic"), a)
	}

	var isErr = make(chan error)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil {
			isErr <- err
		}
	}()

	select {
	case err := <-isErr:
		server.logger.Error("Server", fmt.Sprintf("error while starting %v-backend server", server.appConfig.AppName), err)
	case <-ticker.C:
		server.logger.Info("Server", fmt.Sprintf("%v-backend server started at %v", server.appConfig.AppName, server.httpServer.Addr), nil)
		break
	}
}

func (server *ginServer) GracefulShutdown() {
	listenToSignalNotification()

	if err := server.httpServer.Shutdown(context.Background()); err != nil {
		server.logger.Error("Server", "Server shutdown failed, Forcing server shutdown...", err)

	} else {
		server.logger.Info("Server", "Server graceful shutdown successful.", nil)
	}
}

// fast http  server

func newFastHttpServer(configs *config.Config, api inbound.APIPort, logger logger.Logger) *fasthttpServer {

	return &fasthttpServer{
		appConfig: configs,
		logger:    logger,
		api:       api,
	}
}

func (server *fasthttpServer) Start() {
	port := fmt.Sprintf(":%v", server.appConfig.AppConfig.AppPort)
	server.httpServer = &fasthttp.Server{
		Handler: fasthttproutes.NewRouter(server.appConfig, server.api, server.logger).SetRouters(),
	}

	go func() {
		fmt.Printf("\"regulator server started at %v\n", port)
		if err := server.httpServer.ListenAndServe(port); err != nil {

			server.logger.Error("ServerModule", "Server failure error:", err)
		}
	}()
}

// TODO: resolve commented code
func (server *fasthttpServer) GracefulShutdown() {
	listenToSignalNotification()

	if err := server.httpServer.Shutdown(); err != nil {
		server.logger.Error("ServerModule", "Server shutdown failed, Forcing server shutdown...", err)
	} else {
		server.logger.Info("ServerModule", "Server graceful shutdown successful.", nil)
	}
}

func listenToSignalNotification() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
