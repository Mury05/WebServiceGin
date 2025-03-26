package migrations

import (
	"example/web-service-gin/models"
	"log"

	"gorm.io/gorm"
)

// Migrations automatique pour créer la table albums
func MigrateAlbums(db gorm.DB) {
	err := db.AutoMigrate(&models.Album{})
	if err != nil {
		log.Fatal("Erreur lors de la migration :", err)
	}
}
