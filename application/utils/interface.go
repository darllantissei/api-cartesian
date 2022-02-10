package utils

import "github.com/darllantissei/api-cartesian/application/models"

type IUtilsService interface {
	BuildLogRequestsReceived(params models.ParamsBuildLogsRequestsReceived)
	ParseToSHA256(value string) string
}