package internal

import (
	"net/http"

	"github.com/ShreyashSri/test/internal/boot"
	"github.com/ShreyashSri/test/internal/middleware"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// Public routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Seeqlo is running!",
		})
	})

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.FirebaseAuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			uid, _ := c.Get("uid")
			c.JSON(http.StatusOK, gin.H{
				"message": "Authenticated",
				"uid":     uid,
			})
		})
	}

	port := boot.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
