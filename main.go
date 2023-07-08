package main

import (
	"github.com/giavudangle/go-rest-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)

	router.Run("localhost:8080")

}
