package router

import (
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine, handlers ports.ApiHandler) {
	bikesRouter := router.Group("/bikes")

	bikesRouter.GET("/search", handlers.GetAllBikesHandler)
}
