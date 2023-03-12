package controller

import (
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterAccountRequest struct {
	PublicKey string `json:"publicKey"`
	UserId    string `json:"userId"`
	Nickname  string `json:"nickname"`
}

func (c *Controller) RegisterAccount(ctx *gin.Context) {
	var registerAccountReq RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&registerAccountReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if registerAccountReq.PublicKey == "" || registerAccountReq.UserId == "" || registerAccountReq.Nickname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "publicKey and userId and nickname are required"})
		return
	}

	if registerAccountReq.UserId != ctx.GetString("id") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	user, err := c.Database.GetUserById(registerAccountReq.UserId)
	if err != nil || user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	searchedPublicKey, err := c.Database.GetAccountByPublicKey(registerAccountReq.PublicKey)
	if searchedPublicKey != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "publicKey already exists"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account := repository.Account{
		PublicKey: registerAccountReq.PublicKey,
		UserId:    registerAccountReq.UserId,
		Nickname:  registerAccountReq.Nickname,
	}

	accountCreated, err := c.Database.AddAccount(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, accountCreated)
}
