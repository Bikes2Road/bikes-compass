package ports

import "github.com/gin-gonic/gin"

type ApiHandler interface {
	GetAllBikesHandler(g *gin.Context)
}
