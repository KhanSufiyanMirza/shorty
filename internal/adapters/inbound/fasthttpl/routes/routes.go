package routes

import (
	"hex/internal/adapters/inbound/fasthttpl/handlers/urlshortener"
	"hex/utils/logger"

	fasthttprouter "github.com/fasthttp/router"
)

// User Onboarding Router
func (r *router) routes(log logger.Logger) {

	rg := r.router.Group("/v1/shorty")
	{
		{
			r.urlShortenerRoutes(rg)
		}
	}

}

func (r router) urlShortenerRoutes(rg *fasthttprouter.Group) {
	handler := urlshortener.NewUrlShortenerHandler(r.api, r.logger)

	{
		rg.GET("/{url}",
			handler.Get)

		rg.POST("/",
			handler.Save)

	}
}
