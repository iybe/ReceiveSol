package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	var actualUser LoginUserRequest
	if err := ctx.ShouldBindJSON(&actualUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if actualUser.Username == "" || actualUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	user, err := c.Database.GetUser(actualUser.Username)
	if err != nil || user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	response, err := c.SSOExternal.CreateToken(user.ID, actualUser.Password)
	if err != nil || response == nil || response.Token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	loginUserResponse := LoginUserResponse{
		Token: response.Token,
		Id:    user.ID,
	}

	ctx.JSON(http.StatusOK, loginUserResponse)
}
