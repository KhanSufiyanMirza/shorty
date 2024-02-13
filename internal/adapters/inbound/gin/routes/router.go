package routes

import (
	"hex/config"
	"hex/internal/adapters/inbound/gin/middlewares"

	"hex/internal/ports/inbound"
	"hex/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	// routerConfig     *RouterConfig
	router    *gin.Engine
	appConfig config.Config
	logger    logger.Logger
	api       inbound.APIPort
}

func NewRouter(configs config.Config, api inbound.APIPort, logger logger.Logger) *router {
	return &router{
		// routerConfig:     NewConfig().SetTimeout(configs.GetAppConfig().Timeout),
		router:    gin.New(),
		appConfig: configs,
		logger:    logger,
		api:       api,
	}
}

func (r *router) SetRouters() http.Handler {

	r.routes(r.logger)

	return r.router.Handler()
}

func attachDefaultmiddlewares(r *gin.Engine, log logger.Logger) {
	r.Use(gin.Recovery())
	// r.Use(middlewares.AccessLogger(log))
	r.Use(middlewares.SecurityHeaders())
	r.Use(middlewares.Cors())
	r.Use(middlewares.RateLimit(log))
	r.Use(middlewares.SanitizeQueryMap(log))
}
