package sso

import "backend/service"

const (
	methodCreateUser  = "CreateUser"
	methodCreateToken = "CreateToken"
	methodVerifyToken = "VerifyToken"
)

type errorResponse struct {
	Error string `json:"error"`
}

type Client struct {
	URL           string
	Log           *service.ExternalLog
	Authorization string
}
