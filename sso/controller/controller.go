package controller

import "sso/repository"

type Controller struct {
	Database               *repository.ClientMongoDB
	PasswordSecret         string
	TokenSecret            string
	ExpirationSecondsToken int
}
