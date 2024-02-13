package routes

import (
	"hex/config"

	"hex/internal/ports/inbound"
	"hex/utils/logger"

	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type router struct {
	router    *fasthttprouter.Router
	appConfig *config.Config
	logger    logger.Logger
	api       inbound.APIPort
}

func NewRouter(configs *config.Config, api inbound.APIPort, logger logger.Logger) *router {
	return &router{
		router:    fasthttprouter.New(),
		appConfig: configs,
		logger:    logger,
		api:       api,
	}
}

func (r *router) SetRouters() fasthttp.RequestHandler {

	r.routes(r.logger)

	return r.router.Handler
}
