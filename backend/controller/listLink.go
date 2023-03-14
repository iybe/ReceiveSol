package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ListLink(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing userId"})
		return
	}

	if userId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	network := ctx.Query("network")
	recipient := ctx.Query("recipient")
	status := ctx.Query("status")
	permalink := ctx.Query("permalink")
	permalinkB := false
	if permalink != "" {
		permalinkB = true
	}

	links, err := c.Database.ListLink(userId, status, network, recipient, permalinkB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, links)
}
