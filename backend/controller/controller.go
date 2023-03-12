package controller

import (
	"backend/external/monitor"
	"backend/external/sso"
	"backend/repository"
)

type Controller struct {
	Database        *repository.ClientMongoDB
	SSOExternal     *sso.Client
	MonitorExternal *monitor.Client
}
