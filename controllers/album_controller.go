package controllers

import (
	"net/http"

	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// Obtenir la liste des albums en format json
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

// Récupérer un album spécifique à partir de l'id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// addAlbum pour ajouter un album sous forme de json via la request body
func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Bind JSON and check for errors
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Collecter les messages d'erreur de validation
	var validationErrors []string

	if newAlbum.ID == "" {
		validationErrors = append(validationErrors, "L'ID de l'album est requis")
	}
	if newAlbum.Title == "" {
		validationErrors = append(validationErrors, "Le titre de l'album est requis")
	}
	if newAlbum.Artist == "" {
		validationErrors = append(validationErrors, "L'artiste de l'album est requis")
	}
	if newAlbum.Price <= 0 {
		validationErrors = append(validationErrors, "Le prix de l'album doit être supérieur à 0")
	}

	// Vérifier l'unicité de l'ID
	for _, a := range models.Albums {
		if a.ID == newAlbum.ID {
			validationErrors = append(validationErrors, "Un album avec cet ID existe déjà")
			break
		}
	}

	// Si des erreurs existent, les retourner
	if len(validationErrors) > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	// Ajouter le nouvel album au slice et retourner le résultat
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
