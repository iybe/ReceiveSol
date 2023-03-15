package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserPermaLinkResponse struct {
	RecipientPermaLink string `json:"recipientPermaLink"`
	NetworkPermaLink   string `json:"networkPermaLink"`
	Url                string `json:"url"`
}

func (c *Controller) GetUserPermaLink(ctx *gin.Context) {
	userId := ctx.Query("userId")

	if userId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	user, err := c.Database.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	if user.NetworkPermaLink == "" || user.RecipientPermaLink == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user has not set perma links"})
		return
	}

	getUserPermaLinkResponse := GetUserPermaLinkResponse{
		RecipientPermaLink: user.RecipientPermaLink,
		NetworkPermaLink:   user.NetworkPermaLink,
		Url:                fmt.Sprintf("%s/permalink/%s", c.Url, userId),
	}

	ctx.JSON(http.StatusOK, getUserPermaLinkResponse)
}
