package controller

import (
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var newUser CreateUserRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	userSearched, err := c.Database.GetUser(newUser.Username)
	if userSearched != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRepo := repository.User{
		Username: newUser.Username,
	}

	userCreated, err := c.Database.AddUser(userRepo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := c.SSOExternal.CreateUser(userCreated.ID, newUser.Password)
	if err != nil || response == nil || response.Id == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
