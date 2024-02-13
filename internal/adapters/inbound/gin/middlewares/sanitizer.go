package middlewares

import (
	"hex/models/dto"
	"hex/utils/logger"
	sanitizeUtils "hex/utils/sanitize"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SanitizeQueryMap(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, value := range ctx.Request.URL.Query() {
			for _, v := range value {
				err := sanitizeUtils.Sanitize(v)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest,
						dto.NewHttpErrorMsg(false, http.StatusBadRequest, "invalid query parameters"))
					return
				}
			}
		}

		for _, value := range ctx.Params {
			err := sanitizeUtils.Sanitize(value.Value)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest,
					dto.NewHttpErrorMsg(false, http.StatusBadRequest, "invalid path parameters"))
				return
			}
		}
	}
}
