package handlers

import "github.com/valyala/fasthttp"

//TODO: should be in ports?

type BaseHandler interface {
	Get(ctx *fasthttp.RequestCtx)
	GetAll(ctx *fasthttp.RequestCtx)
	Save(ctx *fasthttp.RequestCtx)
	Update(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
	Options(ctx *fasthttp.RequestCtx)
}
