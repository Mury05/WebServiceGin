package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/migrations"
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	migrations.MigrateAlbums(*database.DB)

	router := gin.Default()
	routes.AlbumRoutes(router)

	router.Run("localhost:8000")
}
