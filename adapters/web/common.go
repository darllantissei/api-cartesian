package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darllantissei/api-cartesian/application/models"
	statusapplication "github.com/darllantissei/api-cartesian/application/status_application"
	"github.com/labstack/echo"
)

func (w *WebServer) errorServer(err error, ectx echo.Context) {
	var (
		code      = http.StatusInternalServerError
		errResult models.Returns
	)

	errResult.Return.Status = statusapplication.Error

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		errResult.Return.Message = append(errResult.Return.Message, fmt.Sprintf("%#v", he.Message))
		if he.Internal != nil {
			errResult.Return.Message = append(errResult.Return.Message, fmt.Sprintf("%v, %v", err, he.Internal))
		}
	} else {
		errResult.Return.Message = append(errResult.Return.Message, http.StatusText(code))
	}

	if !ectx.Response().Committed {
		if ectx.Request().Method == echo.HEAD {
			err = ectx.NoContent(code)
		} else {
			err = ectx.JSON(code, errResult)
		}
		if err != nil {
			ectx.Echo().Logger.Error(err)
		}
	} else {
		log.Print(err.Error())
	}
}
