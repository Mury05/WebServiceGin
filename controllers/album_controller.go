package controllers

import (
	"net/http"

	"example/web-service-gin/database"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Obtenir la liste des albums en format json
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	if err := database.DB.Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// Récupérer un album spécifique à partir de l'id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var albumFind models.Album

	err := database.DB.First(&albumFind, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "erreur serveur"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, albumFind)
}

// addAlbum pour ajouter un album sous forme de json via la request body
func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Bind JSON and check for errors
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Collecter les erreurs de validation dans une map
	validationErrors := make(map[string]string)

	if newAlbum.Title == "" {
		validationErrors["title"] = "Le titre de l'album est requis"
	}
	if newAlbum.Artist == "" {
		validationErrors["artist"] = "L'artiste de l'album est requis"
	}
	if newAlbum.Price <= 0 {
		validationErrors["price"] = "Le prix de l'album doit être supérieur à 0"
	}

	// Si des erreurs existent, les retourner
	if len(validationErrors) > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	// Sauvegarder l'album dans la base de données
	if err := database.DB.Create(&newAlbum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ajout de l'album"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func EditAlbum(c *gin.Context) {
	id := c.Param("id")

	var albumToEdit models.Album

	// Vérifier si l'album existe
	err := database.DB.First(&albumToEdit, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "erreur serveur"})
		}
		return
	}

	// Lire les données de la requête
	var input map[string]interface{}
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Collecter les erreurs de validation
	validationErrors := make(map[string]string)

	if title, ok := input["title"].(string); ok && title == "" {
		validationErrors["title"] = "Le titre de l'album est requis"
	}
	if artist, ok := input["artist"].(string); ok && artist == "" {
		validationErrors["artist"] = "L'artiste de l'album est requis"
	}
	if price, ok := input["price"].(float64); ok && price <= 0 {
		validationErrors["price"] = "Le prix de l'album doit être supérieur à 0"
	}

	// Retourner les erreurs s'il y en a
	if len(validationErrors) > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	// Mise à jour de l'album uniquement avec les champs fournis
	if err := database.DB.Model(&albumToEdit).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'album"})
		return
	}

	// Retourner l'album mis à jour
	c.IndentedJSON(http.StatusOK, albumToEdit)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	var albumToDelete models.Album

	// Vérifier si l'album existe
	err := database.DB.First(&albumToDelete, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "erreur serveur"})
		}
		return
	}

	if err := database.DB.Delete(&albumToDelete).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'album"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Album supprimé avec succès !"})

}
