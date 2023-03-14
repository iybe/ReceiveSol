package controller

import (
	"backend/external/monitor"
	"backend/external/sso"
	"backend/repository"
)

type Controller struct {
	Url             string
	Database        *repository.ClientMongoDB
	SSOExternal     *sso.Client
	MonitorExternal *monitor.Client
}
