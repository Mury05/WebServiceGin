package routes

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	router.Use(middlewares.Logger())
	router.Use(middlewares.RequestLogger())
	router.Use(middlewares.Recovery())
	router.Use(middlewares.RateLimiter())
	api := router.Group("/api")
	{
		api.GET("albums", controllers.GetAlbums)
		api.POST("albums", controllers.AddAlbum)
		api.GET("albums/:id", controllers.GetAlbumByID)
		api.PUT("albums/:id", controllers.EditAlbum)
		api.DELETE("albums/:id", controllers.DeleteAlbum)
	}
}
