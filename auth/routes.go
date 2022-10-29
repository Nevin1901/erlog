package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Secret = []byte("Secret")

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("key")

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Next()
}

func Login( c*gin.Context) {
}