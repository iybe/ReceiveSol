package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Client) AuthMiddleware(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	token := ctx.GetHeader("token")

	if id == "" || token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing headers id or token"})
		return
	}

	err := c.SSOExternal.VerifyToken(id, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	ctx.Set("id", id)

	ctx.Next()
}
