package models

import (
	"strings"

	statusapplication "github.com/darllantissei/api-cartesian/application/status_application"
)

type Returns struct {
	Return `json:"return" valid:"-"`
}

type Return struct {
	Status  statusapplication.StatusApp `json:"status" valid:"-"`  // Status a ser retornado
	Message []string                    `json:"message" valid:"-"` // mensagens a serem retornadas
}

func (r Returns) Error() string {
	return strings.Join(r.Return.Message, ";")
}
