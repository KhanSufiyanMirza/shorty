package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("X-Frame-Options", "deny")
		ctx.Writer.Header().Set("Content-Security-Policy", "self")
		ctx.Writer.Header().Set("Referrer-Policy", "no-referrer")
		ctx.Writer.Header().Set("Server", "")
		ctx.Writer.Header().Set("X-Powered-By", "")
		ctx.Next()
	}
}
