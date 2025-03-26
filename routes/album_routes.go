package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("albums", controllers.GetAlbums)
		api.GET("albums/:id", controllers.GetAlbumByID)
		api.POST("albums", controllers.AddAlbum)
	}
}
