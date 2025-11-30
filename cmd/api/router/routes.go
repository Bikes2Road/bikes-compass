package router

import (
	"github.com/Bikes2Road/bikes-compass/docs"
	"github.com/Bikes2Road/bikes-compass/pkg/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers(router *gin.Engine, handlers ports.ApiHandler, port string) {
	// Configurar el host de Swagger dinámicamente basado en el puerto
	docs.SwaggerInfo.Host = "localhost:" + port

	v1 := router.Group("/v1")
	v1.GET("/health", handlers.HealthHandler)

	bikesRouter := v1.Group("/bikes")

	bikesRouter.GET("/byke/:hash_byke", handlers.GetBykeHandler)
	bikesRouter.GET("/search", handlers.GetAllBikesHandler)
	bikesRouter.GET("/placeholder", handlers.PlaceHolderHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
