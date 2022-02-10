package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/eucatur/go-toolbox/log"
)

type UtilsService struct{}

func (u *UtilsService) ParseToSHA256(value string) string {
	cryptsha256 := sha256.New()
	cryptsha256.Write([]byte(value))
	returnSHA256 := fmt.Sprintf("%x", cryptsha256.Sum(nil))

	return returnSHA256
}

func (u *UtilsService) BuildLogRequestsReceived(params models.ParamsBuildLogsRequestsReceived) {
	go func(prm models.ParamsBuildLogsRequestsReceived) {

		logBody := struct {
			IP             string      `json:"ip"`
			Authorization  string      `json:"authorization"`
			Method         string      `json:"method"`
			URL            string      `json:"url"`
			RequestPayload interface{} `json:"request_payload"`
			Status         int         `json:"status"`
			Response       interface{} `json:"response"`
		}{
			IP:             prm.Ectx.RealIP(),
			Authorization:  prm.Ectx.Request().Header.Get("Authorization"),
			Method:         prm.Ectx.Request().Method,
			URL:            prm.Ectx.Request().RequestURI,
			RequestPayload: string(prm.ReqBody),
			Status:         prm.Ectx.Response().Status,
			Response:       string(prm.ResBody),
		}

		logBytes, err := json.Marshal(logBody)
		if err != nil {

			log.Debug("Error to parse json content log: ", err)
			return
		}

		err = log.File(time.Now().Format(prm.PathLogs), string(logBytes))
		if err != nil {
			log.Debug("Error to generate log file: ", err)
		}

	}(params)
}
