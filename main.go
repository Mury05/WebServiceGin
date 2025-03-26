package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 23.99},
	{ID: "5", Title: "Kind of Blue", Artist: "Miles Davis", Price: 19.99},
	{ID: "6", Title: "A Love Supreme", Artist: "John Coltrane", Price: 29.99},
}

func main() {
	router := gin.Default()
	router.GET("albums", getAlbums)
	router.POST("albums", addAlbum)
	router.Run("localhost:8000")
}

// Obtenir la liste des albums en format json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// addAlbum pour ajouter un album sous forme de json via la request body
func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Ajouter le nouvel album au slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
