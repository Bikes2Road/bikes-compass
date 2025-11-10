package main

import (
	"log"

	"github.com/Bikes2Road/bikes-compass/cmd/api/dependencies"
	_ "github.com/Bikes2Road/bikes-compass/docs"
	"github.com/Bikes2Road/bikes-compass/utils/env"
	"github.com/gin-gonic/gin"
)

// @title           Bikes 2 Road API
// @version         1.0
// @description     This is the docs of Bikes 2 Road API.

// @host      localhost:8081
// @BasePath  /v1/bikes

func main() {
	server := gin.Default()

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
