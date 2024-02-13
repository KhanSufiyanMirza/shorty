package urlshortener

//TODO: move to handlers package instead of wallet like issuance
import (
	"errors"
	"hex/constants"
	"hex/internal/adapters/inbound/gin/handlers"
	"hex/internal/ports/inbound"
	"hex/models/dto"
	utilshttp "hex/utils/http"
	"hex/utils/logger"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
func (handler *handler) Save(c *gin.Context) {
	var req dto.CustomShortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utilshttp.StatusBadRequest(c.Writer, c.Request, err)
		return
	}
	// check if the input is an actual URL
	if !govalidator.IsURL(req.URL) {
		utilshttp.StatusBadRequest(c.Writer, c.Request, constants.ErrInvalidUrl)
		return
	}
	resp, err := handler.api.GetUrlShortingService().CreateShorty(c, &req, c.ClientIP())
	if err != nil {
		utilshttp.StatusServiceUnavailableError(c.Writer, c.Request, err)
		return
	}
	utilshttp.StatusOK(c.Writer, c.Request, resp)
}

func (handler *handler) Get(c *gin.Context) {
	customUrl := c.Param("url")
	url, err := handler.api.GetUrlShortingService().ResolveShorty(c, customUrl)
	if err != nil {
		if errors.Is(err, constants.ErrRecordNotFound) {
			utilshttp.StatusNotFound(c.Writer, c.Request, err)
		} else {
			utilshttp.StatusInternalServerError(c.Writer, c.Request, err)
		}
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}
