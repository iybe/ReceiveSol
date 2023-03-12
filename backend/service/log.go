package service

import "log"

type Log struct {
	ExternalLog *ExternalLog
}

type ExternalLog struct {
	HostName string
}

func (e *ExternalLog) Error(method string, err error) {
	log.Printf("Error - external: %s - metodo: %s - error: %v", e.HostName, method, err)
}

func (e *ExternalLog) WarnWithStatusCode(method string, err error, statusCode int) {
	log.Printf("Warn - external: %s - code: %d - metodo: %s - error: %v", e.HostName, statusCode, method, err)
}
