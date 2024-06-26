// pkg/middleware/auth.go
package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-architecture/pkg/services"
)

func BasicAuthMiddleware(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		token := authParts[1]
		decodedToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(decodedToken), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
			c.Abort()
			return
		}

		name := credentials[0]
		password := credentials[1]

		// Authenticate user
		authenticatedUser, err := userService.AuthenticateUser(name, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			c.Abort()
			return
		}

		// Store user information in context if needed
		c.Set("user", authenticatedUser)

		// Proceed to the next handler if authentication is successful
		c.Next()
	}
}
