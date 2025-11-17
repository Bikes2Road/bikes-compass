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

// Get All Bikes
// @Summary Search Bikes
// @Description This service extract all bikes with pagination, you can search bikes by name, or all bikes
// @Tags Bikes 2 Road
// @Param page query int false "page that you want extract" minimum(1)
// @Param cant query int false "cant bikes you want extract" maximum(30)
// @Param name query string false "name of byke that you want search" example(BMW M1000RR)
// @Param brand query string false "brand of byke that you want search" example(BMW)
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
		matched, _ := regexp.MatchString(`^[A-Za-z0-9\s]+$`, queryRequest.Name)
		if !matched {
			errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidStringBike, nil)
			c.JSON(errResponse.Code, errResponse)
			return
		}
	}

	if queryRequest.Brand != "" {
		matched, _ := regexp.MatchString(`^[A-Za-z\s]+$`, queryRequest.Brand)
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

// Get Byke
// @Summary Search Byke by Hash
// @Description This service extract all data from a Byke by Hash_Byke
// @Tags Bikes 2 Road
// @Param hash_byke path string true "Hash of Byke that you want extract"
// @Produce json
// @Success 200 {object} domain.GetBykeResponseSuccess
// @Failure 400 {object} domain.ResponseHttpError
// @Failure 404 {object} domain.ResponseHttpError
// @Failure 401 {object} domain.ResponseHttpError
// @Failure 500 {object} domain.ResponseHttpError
// @Router /byke/{hash_byke} [get]
func (h *ApiHandler) GetBykeHandler(c *gin.Context) {
	var paramRequest domain.SearchBykeRequest

	pathRequest := c.Request.RequestURI

	if err := c.ShouldBindUri(&paramRequest); err != nil {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidPathParams, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	// INSERT_YOUR_CODE
	// Validar que HashByke sea un string alfanumérico de exactamente 12 dígitos usando regex
	matched, _ := regexp.MatchString(`^[A-Za-z0-9]{12}$`, paramRequest.HashByke)
	if !matched {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidPathParam, nil)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	byke, errResp := h.application.GetByke.Execute(h.ctx, paramRequest, pathRequest)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, byke)
}

// Placeholder
// @Summary Search Byke by Hash
// @Description This service extract a list of name bikes from a Byke by name
// @Tags Bikes 2 Road
// @Param name query string false "name of byke that you want search" example(Yamaha)
// @Produce json
// @Success 200 {object} domain.PlaceHolderResponseSuccess
// @Failure 400 {object} domain.ResponseHttpError
// @Failure 404 {object} domain.ResponseHttpError
// @Failure 401 {object} domain.ResponseHttpError
// @Failure 500 {object} domain.ResponseHttpError
// @Router /placeholder [get]
func (h *ApiHandler) PlaceHolderHandler(c *gin.Context) {
	var queryRequest domain.PlaceHolderRequest

	err := c.BindQuery(&queryRequest)
	if err != nil {
		errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidQueryParams, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	if queryRequest.NameByke != "" {
		matched, _ := regexp.MatchString(`^[A-Za-z0-9\s]+$`, queryRequest.NameByke)
		if !matched {
			errResponse := errorBikes.MapErrorResponse(errorBikes.ErrorInvalidStringBike, nil)
			c.JSON(errResponse.Code, errResponse)
			return
		}
	}

	bikes, errResp := h.application.PlaceHolder.Execute(h.ctx, queryRequest)
	if errResp != nil {
		c.JSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, bikes)

}
