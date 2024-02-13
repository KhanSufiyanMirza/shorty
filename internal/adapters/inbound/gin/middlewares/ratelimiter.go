package middlewares

import (
	"hex/models/dto"
	"hex/utils/logger"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const MaxNoOfRequestAllowed = 50

func RateLimit(log logger.Logger) gin.HandlerFunc {
	allowedRequest, err := strconv.Atoi(os.Getenv("ENV_MAX_ALLOWED_REQUEST"))
	if err != nil {
		allowedRequest = MaxNoOfRequestAllowed
	}

	limiter := rate.NewLimiter(per(time.Second, allowedRequest), allowedRequest)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				dto.NewHttpErrorMsg(false, http.StatusTooManyRequests, "too Many Requests"))
			return
		}
		c.Next()
	}
}

func per(duration time.Duration, eventCount int) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}
