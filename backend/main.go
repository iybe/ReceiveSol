package main

import (
	"backend/controller"
	"backend/external/monitor"
	"backend/external/sso"
	"backend/middleware"
	"backend/repository"
	"backend/service"
	"context"
	"log"
	"os"

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
	MONGODB_COLLECTION_ACCOUNT := os.Getenv("MONGODB_COLLECTION_ACCOUNT")
	MONGODB_COLLECTION_LINK := os.Getenv("MONGODB_COLLECTION_LINK")
	EXTERNAL_SSO_URL := os.Getenv("EXTERNAL_SSO_URL")
	EXTERNAL_SSO_HOSTNAME := os.Getenv("EXTERNAL_SSO_HOSTNAME")
	EXTERNAL_SSO_AUTHORIZATION := os.Getenv("EXTERNAL_SSO_AUTHORIZATION")
	EXTERNAL_MONITOR_URL := os.Getenv("EXTERNAL_MONITOR_URL")
	EXTERNAL_MONITOR_HOSTNAME := os.Getenv("EXTERNAL_MONITOR_HOSTNAME")
	EXTERNAL_MONITOR_AUTHORIZATION := os.Getenv("EXTERNAL_MONITOR_AUTHORIZATION")

	client, err := repository.CreateClient(MONGODB_URI)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Disconnect(context.Background())

	mongoClient := repository.ClientMongoDB{
		Client:            client,
		DatabaseName:      MONGODB_DATABASE,
		CollectionUsers:   MONGODB_COLLECTION_USER,
		CollectionAccount: MONGODB_COLLECTION_ACCOUNT,
		CollectionLink:    MONGODB_COLLECTION_LINK,
	}

	ssoClient := sso.Client{
		URL: EXTERNAL_SSO_URL,
		Log: &service.ExternalLog{
			HostName: EXTERNAL_SSO_HOSTNAME,
		},
		Authorization: EXTERNAL_SSO_AUTHORIZATION,
	}

	monitorClient := monitor.Client{
		URL: EXTERNAL_MONITOR_URL,
		Log: &service.ExternalLog{
			HostName: EXTERNAL_MONITOR_HOSTNAME,
		},
		Authorization: EXTERNAL_MONITOR_AUTHORIZATION,
	}

	middlewareClient := middleware.Client{
		SSOExternal: &ssoClient,
	}

	controllerClient := controller.Controller{
		Database:        &mongoClient,
		SSOExternal:     &ssoClient,
		MonitorExternal: &monitorClient,
	}

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.POST("/user", controllerClient.CreateUser)
	router.POST("/user/login", controllerClient.LoginUser)
	router.POST("/user/permalink", middlewareClient.AuthMiddleware, controllerClient.UpdateUserPermaLink)

	router.POST("/account", middlewareClient.AuthMiddleware, controllerClient.RegisterAccount)
	router.GET("/account", middlewareClient.AuthMiddleware, controllerClient.ListAccount)

	router.POST("/link", middlewareClient.AuthMiddleware, controllerClient.CreateLink)
	router.GET("/link", middlewareClient.AuthMiddleware, controllerClient.ListLink)

	router.GET("/solanalink", controllerClient.GetLink)

	router.POST("/permalink", controllerClient.CreatePermaLink)

	log.Fatal(router.Run(":" + PORT))
}
