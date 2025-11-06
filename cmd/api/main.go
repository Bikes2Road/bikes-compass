package main

import (
	"log"

	"github.com/Bikes2Road/bikes-compass/cmd/api/dependencies"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	wrappers := dependencies.DefaultWrappers()

	err := dependencies.InitDependencies(wrappers, server)
	if err != nil {
		log.Fatalf("Error inicializando dependencias: %v", err)
	}

	log.Println("Servidor iniciando en puerto 8080...")
	if err := server.Run(":8081"); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
