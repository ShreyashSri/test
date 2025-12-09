package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ShreyashSri/test/internal/boot"
	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			c.Abort()
			return
		}

		client, err := boot.App.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Firebase Auth client error"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("uid", token.UID)
		c.Set("claims", token.Claims)
		c.Next()
	}
}
