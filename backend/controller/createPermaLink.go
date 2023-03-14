package controller

import (
	"backend/repository"
	"backend/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePermaLinkRequest struct {
	UserId         string  `json:"userId"`
	ExpectedAmount float64 `json:"expectedAmount"`
}

func (c *Controller) CreatePermaLink(ctx *gin.Context) {
	var newPermaLink CreatePermaLinkRequest
	if err := ctx.ShouldBindJSON(&newPermaLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newPermaLink.ExpectedAmount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "expected amount must be greater than 0"})
		return
	}

	user, err := c.Database.GetUserById(newPermaLink.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if user.RecipientPermaLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user must have a recipient perma link"})
		return
	}

	account, err := c.Database.GetAccount(user.RecipientPermaLink, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "recipient must be a public key of the user"})
		return
	}

	reference, err := service.GenerateRandomPublicKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	solanaPayLink := service.CreateSolanaPayLink(user.RecipientPermaLink, reference, newPermaLink.ExpectedAmount)

	linkRepo := repository.Link{
		UserId:         user.ID,
		AccountId:      account.ID,
		Link:           solanaPayLink,
		Reference:      reference,
		Recipient:      user.RecipientPermaLink,
		Network:        user.NetworkPermaLink,
		ExpectedAmount: newPermaLink.ExpectedAmount,
		Status:         LinkStatusCreated,
		Expired:        false,
		IsPermaLink:    true,
	}

	linkCreated, err := c.Database.CreateLink(linkRepo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()

	err = c.MonitorExternal.RegisterLink(linkCreated.ID, linkCreated.Reference, linkCreated.Recipient, linkCreated.Network, linkCreated.ExpectedAmount, 0, now)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.Database.UpdateLinkStatus(linkCreated.ID, LinkStatusPending)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createLinkResponse := CreateLinkResponse{
		Link:      solanaPayLink,
		Reference: reference,
	}

	ctx.JSON(http.StatusOK, createLinkResponse)
}
