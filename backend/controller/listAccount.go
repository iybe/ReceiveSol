package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ListAccount(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing userId"})
		return
	}

	if userId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	accounts, err := c.Database.ListAccountByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
