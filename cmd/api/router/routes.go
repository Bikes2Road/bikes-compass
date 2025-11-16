package router

import (
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers(router *gin.Engine, handlers ports.ApiHandler) {
	bikesRouter := router.Group("/v1/bikes")

	bikesRouter.GET("/byke/:hash_byke", handlers.GetBykeHandler)
	bikesRouter.GET("/search", handlers.GetAllBikesHandler)
	bikesRouter.GET("/placeholder", handlers.PlaceHolderHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
