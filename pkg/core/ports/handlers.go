package ports

import "github.com/gin-gonic/gin"

type ApiHandler interface {
	GetAllBikesHandler(g *gin.Context)
	GetBykeHandler(g *gin.Context)
	PlaceHolderHandler(g *gin.Context)
}
