package controller

import (
	"net/http"
	"sso/repository"
	"sso/services"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateToken(ctx *gin.Context) {
	var userLogin repository.User
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Database.FindUser(userLogin.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal erro"})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "credencial invalida"})
		return
	}

	if user.Password != services.HashPassword(userLogin.Password, c.PasswordSecret) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "credencial invalida"})
		return
	}

	token, err := services.CreateToken(user.User, c.ExpirationSecondsToken, c.TokenSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal erro"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}
