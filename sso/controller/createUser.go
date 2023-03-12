package controller

import (
	"net/http"
	"sso/repository"
	"sso/services"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	var newUser repository.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.User == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user e password são obrigatórios"})
		return
	}

	user, err := c.Database.FindUser(newUser.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	if user != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "usuario ja existe"})
		return
	}

	newUser.Password = services.HashPassword(newUser.Password, c.PasswordSecret)

	result, err := c.Database.AddUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ocorreu um erro ao criar o usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
