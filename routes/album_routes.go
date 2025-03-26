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
		api.GET("albums/:id", controllers.GetAlbumByID)
		api.POST("albums", controllers.AddAlbum)
	}
}
