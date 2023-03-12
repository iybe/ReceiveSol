package monitor

import "backend/service"

const (
	methodRegisterLink = "registerLink"
)

type errorResponse struct {
	Error string `json:"error"`
}

type Client struct {
	URL           string
	Log           *service.ExternalLog
	Authorization string
}
