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

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for _, a := range models.Albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Un album avec cet id existe déjà"})
		}
	}

	// Ajouter le nouvel album au slice
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
