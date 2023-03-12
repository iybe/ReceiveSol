package controller

import (
	"net/http"
	"sso/services"

	"github.com/gin-gonic/gin"
)

func (c *Controller) VerifyToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	id := ctx.GetHeader("id")

	if token == "" || id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "token e id são obrigatórios"})
		return
	}

	err := services.VerifyToken(token, id, c.TokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
