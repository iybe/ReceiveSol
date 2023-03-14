package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetLink(ctx *gin.Context) {
	reference := ctx.Query("reference")
	if reference == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing reference"})
		return
	}

	link, err := c.Database.SearchByReference(reference)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if link == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	if link.Status != "created" && link.Status != "pending" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	ctx.JSON(http.StatusOK, link)
}
