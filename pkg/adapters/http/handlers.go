package http

import (
	"context"
	"net/http"
	"regexp"

	"github.com/Bikes2Road/bikes-compass/pkg/core"
	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	errorBikes "github.com/Bikes2Road/bikes-compass/utils/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiHandler struct {
	application core.Application
	ctx         context.Context
}

func NewApiHandler(application core.Application) ports.ApiHandler {
	return &ApiHandler{
		application: application,
		ctx:         context.Background(),
	}
}

var validate = validator.New()

// @BasePath /v1/bikes

// Get All Bikes
// @Summary Search Bikes
// @Description This service extract all bikes with pagination, you can search bikes by name, or all bikes
// @Tags Bikes 2 Road
// @Param page query int false "page that you want extract" minimum(1)
// @Param cant query int false "cant bikes you want extract" maximum(30)
// @Param name query string false "name of byke that you want search" example(Yamaha)
// @Produce json
// @Success 200 {object} domain.GetAllResponseSuccess
// @Failure 400 {object} domain.ResponseHttpError
// @Failure 404 {object} domain.ResponseHttpError
// @Failure 401 {object} domain.ResponseHttpError
// @Failure 500 {object} domain.ResponseHttpError
// @Router /search [get]
func (h *ApiHandler) GetAllBikesHandler(c *gin.Context) {
	var queryRequest domain.GetAllBikesRequest

	pathRequest := c.Request.RequestURI

	err := c.BindQuery(&queryRequest)
	if err != nil {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidQueryParams, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	// Set default values
	if queryRequest.Page == 0 {
		queryRequest.Page = 1
	}

	if queryRequest.Cant == 0 {
		queryRequest.Cant = 10
	}

	// Chcck Cant and Page Values
	if queryRequest.Page <= 0 {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidPage, nil)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	if queryRequest.Cant > 30 {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidCant, nil)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	// INSERT_YOUR_CODE
	// Validar que Name solo contenga letras (mayúsculas, minúsculas, espacios) o esté vacío usando regex
	if queryRequest.Name != "" {
		matched, _ := regexp.MatchString(`^[A-Za-z\s]+$`, queryRequest.Name)
		if !matched {
			errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidStringBike, nil)
			c.JSON(errResponse.Code, errResponse)
			return
		}
	}

	bikes, errResp := h.application.GetAllBikes.Execute(h.ctx, queryRequest, pathRequest)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, bikes)
}
