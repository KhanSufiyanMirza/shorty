package routes

import (
	"hex/internal/adapters/inbound/gin/handlers/urlshortener"
	"hex/utils/logger"

	"github.com/gin-gonic/gin"
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

func (r router) urlShortenerRoutes(rg *gin.RouterGroup) {
	handler := urlshortener.NewUrlShortenerHandler(r.api, r.logger)

	{
		rg.GET("/:url",
			handler.Get)

		rg.POST("/",
			handler.Save)

	}
}
