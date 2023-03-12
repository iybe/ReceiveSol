package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Client) AuthMiddleware(ctx *gin.Context) {
	authorization := ctx.GetHeader("authorization")

	if authorization == "" || authorization != c.AuthorizationSecret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	ctx.Next()
}
