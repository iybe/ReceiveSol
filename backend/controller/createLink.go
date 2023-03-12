package controller

import (
	"backend/repository"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	LinkStatusPending = "pending"
	LinkStatusCreated = "created"
)

type CreateLinkRequest struct {
	UserId         string  `json:"userId"`
	Nickname       string  `json:"nickname"`
	Recipient      string  `json:"recipient"`
	Network        string  `json:"network"`
	ExpectedAmount float32 `json:"expectedAmount"`
}

type CreateLinkResponse struct {
	Link      string `json:"link"`
	Reference string `json:"reference"`
}

func (c *Controller) CreateLink(ctx *gin.Context) {
	var newLink CreateLinkRequest
	if err := ctx.ShouldBindJSON(&newLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newLink.UserId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	if newLink.Recipient == "" || newLink.Network == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing recipient or network"})
		return
	}

	if newLink.ExpectedAmount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "expected amount must be greater than 0"})
		return
	}

	if newLink.Network != "mainnet" && newLink.Network != "testnet" && newLink.Network != "devnet" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "network must be mainnet or testnet or devnet"})
		return
	}

	account, err := c.Database.GetAccountByPublicKey(newLink.Recipient)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if account == nil || account.UserId != newLink.UserId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "recipient must be a public key of the user"})
		return
	}

	reference, err := service.GenerateRandomPublicKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	solanaPayLink := service.CreateSolanaPayLink(newLink.Recipient, reference, newLink.ExpectedAmount)

	linkRepo := repository.Link{
		Nickname:       newLink.Nickname,
		UserId:         newLink.UserId,
		AccountId:      account.ID,
		Link:           solanaPayLink,
		Reference:      reference,
		Recipient:      newLink.Recipient,
		Network:        newLink.Network,
		ExpectedAmount: newLink.ExpectedAmount,
		Status:         LinkStatusCreated,
	}

	linkCreated, err := c.Database.CreateLink(linkRepo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.MonitorExternal.RegisterLink(linkCreated.ID, linkCreated.Reference, linkCreated.Recipient, linkCreated.Network, linkCreated.ExpectedAmount)
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

	ctx.JSON(http.StatusCreated, createLinkResponse)
}
