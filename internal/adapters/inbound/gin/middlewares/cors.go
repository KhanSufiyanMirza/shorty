package middlewares

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	allowedDomains := os.Getenv("ENV_ALLOWED_DOMAINS")

	if allowedDomains == "" {
		log.Fatalln("required environment variable ENV_ALLOWED_DOMAINS, but found empty string")
	}

	whitelistedDomains := strings.Split(allowedDomains, ",")

	return cors.New(
		cors.Config{
			AllowAllOrigins:  false,
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Origin", "Cache-Control"},
			AllowOrigins:     whitelistedDomains,
		},
	)
}
