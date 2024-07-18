package middleware

import (
	"os"
	"spotoncars_server/internal"
	"strings"

	"github.com/gin-gonic/gin"
)

func DriverAuthenticationGuard(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

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
	isValid, _, _, clm := internal.Validate(token, true)

	if !isValid {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid token",
		})
		return
	}

	sid, ok := clm[os.Getenv("TOKEN_CLAIM")].(string)
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "SID claim not found",
		})
		return
	}

	c.Set("sid", sid)
	c.Next()
}
