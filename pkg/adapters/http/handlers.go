package http

import (
	"context"
	"net/http"

	"github.com/Bikes2Road/bikes-compass/pkg/core"
	"github.com/Bikes2Road/bikes-compass/pkg/core/domain"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
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

func (h *ApiHandler) GetAllBikesHandler(c *gin.Context) {
	var queryRequest domain.GetAllBikesRequest

	err := c.ShouldBind(&queryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.GetAllResponseError{Status: "Error", Message: err.Error()})
		return
	}

	bikes, err := h.application.GetAllBikes.Execute(h.ctx, queryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.GetAllResponseError{Status: "Error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bikes)
}
