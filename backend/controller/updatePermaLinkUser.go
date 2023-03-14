package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserPermaLinkRequest struct {
	UserId              string `json:"userId"`
	PermaLinkUrl        string `json:"permaLinkUrl"`
	PermaLinkRecipient  string `json:"permaLinkRecipient"`
	PermaLinkExpiration int64  `json:"permaLinkExpiration"`
}

type UpdateUserPermaLinkResponse struct {
	PermaLinkUrl        string `json:"permaLinkUrl"`
	PermaLinkRecipient  string `json:"permaLinkRecipient"`
	PermaLinkExpiration int64  `json:"permaLinkExpiration"`
}

func (c *Controller) UpdateUserPermaLink(ctx *gin.Context) {
	var actualUser UpdateUserPermaLinkRequest
	if err := ctx.ShouldBindJSON(&actualUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if actualUser.UserId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	userUrl, err := c.Database.SearchPermaLinkUrl(actualUser.PermaLinkUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userUrl != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "permaLinkUrl already exists"})
		return
	}

	user, err := c.Database.GetUserById(actualUser.UserId)
	if err != nil || user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = c.Database.UpdateUserPermaLink(actualUser.UserId, actualUser.PermaLinkUrl, actualUser.PermaLinkRecipient, actualUser.PermaLinkExpiration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateUserPermaLinkResponse := UpdateUserPermaLinkResponse{
		PermaLinkUrl:        actualUser.PermaLinkUrl,
		PermaLinkRecipient:  actualUser.PermaLinkRecipient,
		PermaLinkExpiration: actualUser.PermaLinkExpiration,
	}

	ctx.JSON(http.StatusOK, updateUserPermaLinkResponse)
}
