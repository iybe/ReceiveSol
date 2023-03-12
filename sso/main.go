package main

import (
	"context"
	"log"
	"os"
	"sso/controller"
	"sso/middleware"
	"sso/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
		return
	}
}

func main() {
	PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	MONGODB_DATABASE := os.Getenv("MONGODB_DATABASE")
	MONGODB_COLLECTION_USER := os.Getenv("MONGODB_COLLECTION_USER")
	PASSWORD_SECRET := os.Getenv("PASSWORD_SECRET")
	TOKEN_SECRET := os.Getenv("TOKEN_SECRET")
	TOKEN_EXPIRATION_SECONDS, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_SECONDS"))
	if err != nil {
		log.Fatal(err)
		return
	}

	AUTHORIZATION_SECRET := os.Getenv("AUTHORIZATION_SECRET")

	client, err := repository.CreateClient(MONGODB_URI)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Disconnect(context.Background())

	mongoClient := repository.ClientMongoDB{
		Client:          client,
		DatabaseName:    MONGODB_DATABASE,
		CollectionUsers: MONGODB_COLLECTION_USER,
	}

	middlewareClient := middleware.Client{
		AuthorizationSecret: AUTHORIZATION_SECRET,
	}

	controllerClient := controller.Controller{
		Database:               &mongoClient,
		PasswordSecret:         PASSWORD_SECRET,
		TokenSecret:            TOKEN_SECRET,
		ExpirationSecondsToken: TOKEN_EXPIRATION_SECONDS,
	}

	router := gin.Default()

	router.Use(middlewareClient.AuthMiddleware)

	router.POST("/user", controllerClient.CreateUser)
	router.POST("/token", controllerClient.CreateToken)
	router.GET("/token/verify", controllerClient.VerifyToken)

	log.Fatal(router.Run(":" + PORT))
}
