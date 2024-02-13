package handlers

import "github.com/gin-gonic/gin"

//TODO: should be in ports?

type BaseHandler interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Options(ctx *gin.Context)
}
