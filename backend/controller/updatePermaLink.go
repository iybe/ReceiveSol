package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserPermaLinkRequest struct {
	UserId             string `json:"userId"`
	RecipientPermaLink string `json:"recipientPermaLink"`
	NetworkPermaLink   string `json:"networkPermaLink"`
}

func (c *Controller) UpdateUserPermaLink(ctx *gin.Context) {
	var actualUser UpdateUserPermaLinkRequest
	if err := ctx.ShouldBindJSON(&actualUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if actualUser.RecipientPermaLink == "" || actualUser.NetworkPermaLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "perma links must not be empty"})
		return
	}

	if actualUser.UserId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	err := c.Database.SetPermaLink(actualUser.UserId, actualUser.RecipientPermaLink, actualUser.NetworkPermaLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, actualUser)
}
