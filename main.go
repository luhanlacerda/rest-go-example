package main

import (
	albumController "example/rest-go-example/album/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albumController.GetAlbums)
	router.POST("/albums", albumController.PostAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)

	router.Run("localhost:8080")
}
