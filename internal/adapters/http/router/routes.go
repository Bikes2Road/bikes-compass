package router

import (
	"time"

	"github.com/Bikes2Road/bikes-compass/internal/adapters/http/middleware"
	"github.com/Bikes2Road/bikes-compass/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	handlers ports.ApiHandler
}

func NewRouter(handlers ports.ApiHandler) ports.Router {
	return &Router{handlers: handlers}
}

func (r *Router) SetUp(isDevelopment bool) *gin.Engine {
	// Set Gin mode
	if !isDevelopment {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()

	router := gin.New()

	router.Use(middleware.ErrorHandler())
	router.Use(middleware.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Recovery())
	router.Use(middleware.PrometheusMiddleware())

	bikesRouter := router.Group("/api/v1/bikes")
	bikesRouter.GET("/health", r.handlers.HealthHandler)

	bikesRouter.GET("/byke/:hash_byke", r.handlers.GetBykeHandler)
	bikesRouter.GET("/search", r.handlers.GetAllBikesHandler)
	bikesRouter.GET("/placeholder", r.handlers.PlaceHolderHandler)

	bikesRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	bikesRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
