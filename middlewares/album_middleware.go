package middlewares

import (
	"log"
	"net/http"
	"sync"
	"time"

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

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // exécute le prochain middleware/handler
		duration := time.Since(start)
		status := c.Writer.Status()
		log.Printf("%s %s -> %d in %v", c.Request.Method, c.Request.URL.Path, status, duration)
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}

var visitors = make(map[string]time.Time)
var mu sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		visitorIP := c.ClientIP()
		lastSeen, exists := visitors[visitorIP]
		if exists && time.Since(lastSeen) < time.Second {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		visitors[visitorIP] = time.Now()
		c.Next()
	}
}
