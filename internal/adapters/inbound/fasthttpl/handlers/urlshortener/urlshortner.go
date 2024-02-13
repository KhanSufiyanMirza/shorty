package urlshortener

//TODO: move to handlers package instead of wallet like issuance
import (
	"encoding/json"
	"errors"
	"hex/constants"
	"hex/internal/adapters/inbound/fasthttpl/handlers"
	"hex/internal/ports/inbound"
	"hex/models/dto"
	"hex/utils/logger"

	utilsfasthttp "hex/utils/fasthttp"

	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
)

// UrlShortenerHandler is an interface for handling url-shortening-related operations.
type UrlShortenerHandler interface {
	handlers.BaseHandler
}

// handler is a struct that combines functionalities from handlers.BaseHandler,
// an inbound APIPort (api), and a logger.Logger (log)
type handler struct {
	handlers.BaseHandler
	api inbound.APIPort
	log logger.Logger
}

// NewUrlShortenerHandler creates and returns a new instance of the UrlShortenerHandler interface.
func NewUrlShortenerHandler(
	api inbound.APIPort,
	log logger.Logger,

) UrlShortenerHandler {
	return &handler{
		log: log,
		api: api,
	}

}
func (handler *handler) Get(c *fasthttp.RequestCtx) {
	customUrl := c.UserValue("url").(string)
	url, err := handler.api.GetUrlShortingService().ResolveShorty(c, customUrl)
	if err != nil {
		if errors.Is(err, constants.ErrRecordNotFound) {
			utilsfasthttp.StatusNotFound(c, err)
		} else {
			utilsfasthttp.StatusInternalServerError(c, err)
		}
		return
	}
	c.Redirect(url, fasthttp.StatusMovedPermanently)
}

func (handler *handler) Save(c *fasthttp.RequestCtx) {
	var req dto.CustomShortRequest

	if err := json.Unmarshal(c.Request.Body(), &req); err != nil {
		utilsfasthttp.StatusBadRequest(c, err)
		return
	}

	// check if the input is an actual URL
	if !govalidator.IsURL(req.URL) {
		utilsfasthttp.StatusBadRequest(c, constants.ErrInvalidUrl)
		return
	}

	resp, err := handler.api.GetUrlShortingService().CreateShorty(c, &req, c.RemoteIP().String())
	if err != nil {
		utilsfasthttp.StatusServiceUnavailableError(c, err)
		return
	}
	utilsfasthttp.StatusOK(c, resp)
}
