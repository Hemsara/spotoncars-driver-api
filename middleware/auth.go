package middleware

import (
	"fmt"
	"spotoncars_server/internal"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationGuard(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	fmt.Println(authHeader)
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "No authorization header",
		})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid token format",
		})
		return
	}

	token := headerParts[1]
	isValid, _, _ := internal.Validate(token)

	if !isValid {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.Next()
}
