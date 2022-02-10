package models

import (
	"net/http"

	"github.com/labstack/echo"
)

type Utils struct {
	PrivateKey string `json:"-"`
}

type ParamsBuildLogsRequestsReceived struct {
	PathLogs string
	Ectx     echo.Context
	ReqBody  []byte
	ResBody  []byte
}

type ParamsBuildLogsRequestsClients struct {
	PathLogs string
	Response *http.Response
	Sended   []byte
	Body     []byte
}
