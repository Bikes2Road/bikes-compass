package main

import (
	"log"
	"time"

	"github.com/Bikes2Road/bikes-compass/cmd/api/dependencies"
	_ "github.com/Bikes2Road/bikes-compass/docs"
	"github.com/Bikes2Road/bikes-compass/pkg/adapters/http/middleware"
	"github.com/Bikes2Road/bikes-compass/utils/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title           Bikes Compass API
// @version         1.0
// @description     This is the docs of Bikes Compass API from Bikes2Road.

// @BasePath  /v1

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.Use(middleware.PrometheusMiddleware())

	wrappers := dependencies.DefaultWrappers()

	err := dependencies.InitDependencies(wrappers, server)
	if err != nil {
		log.Fatalf("Error inicializando dependencias: %v", err)
	}

	port := env.GetAppPort()

	log.Printf("Servidor iniciando en puerto %s...", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
