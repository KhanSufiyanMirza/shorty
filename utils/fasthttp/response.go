package fasthttp

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

const (
	contentTypeXml  = "application/xml; charset=utf-8"
	contentTypeJson = "application/json"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (resp *response) jsonBytes() []byte {
	data, _ := json.Marshal(resp)
	return data
}

func (resp *response) string() string {
	return string(resp.jsonBytes())
}

func (resp *response) sendResponse(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set("Content-Type", contentTypeJson)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.SetStatusCode(resp.Status)
	// b, _ := xml.Marshal(resp)
	// _, _ = ctx.Write(b)

	ctx.Write(resp.jsonBytes())
	// xml.NewEncoder(ctx.Response.BodyWriter()).Encode(resp)
	// log.Println(resp.string())
}

// 200

func StatusOK(c *fasthttp.RequestCtx, data interface{}) {
	newResponse(data, fasthttp.StatusOK).sendResponse(c)
}

// 204
func StatusNoContent(c *fasthttp.RequestCtx) {
	newResponse(nil, fasthttp.StatusNoContent).sendResponse(c)
}

// 400
func StatusBadRequest(c *fasthttp.RequestCtx, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, fasthttp.StatusBadRequest).sendResponse(c)
}

// 404
func StatusNotFound(c *fasthttp.RequestCtx, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, fasthttp.StatusNotFound).sendResponse(c)
}

// 405
func StatusMethodNotAllowed(c *fasthttp.RequestCtx) {
	newResponse(nil, fasthttp.StatusMethodNotAllowed).sendResponse(c)
}

// 409
func StatusConflict(c *fasthttp.RequestCtx, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, fasthttp.StatusConflict).sendResponse(c)
}

// 500
func StatusInternalServerError(c *fasthttp.RequestCtx, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, fasthttp.StatusInternalServerError).sendResponse(c)
}

// 503
func StatusServiceUnavailableError(c *fasthttp.RequestCtx, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, fasthttp.StatusServiceUnavailable).sendResponse(c)
}
