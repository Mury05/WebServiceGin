package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Vérifie que le chemin complet commence par "/api"
		if len(c.FullPath()) < 4 || c.FullPath()[:4] != "/api" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Route must be prefixed with /api"})
			c.Abort()
			return
		}
		// Continue la chaîne des middlewares et le traitement de la requête
		c.Next()
	}
}
