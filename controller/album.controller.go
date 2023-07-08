package controller

import (
	"net/http"

	"github.com/giavudangle/go-rest-api/domain"
	"github.com/gin-gonic/gin"
)

var albums = []domain.Album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}
