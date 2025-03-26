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
	router.Run("localhost:8000")
}

// Obtenir la liste des albums en format json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
