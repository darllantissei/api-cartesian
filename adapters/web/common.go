package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/darllantissei/api-cartesian/application/common"
	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/eucatur/go-toolbox/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (w *WebServer) Authorization(username, password string, c echo.Context) (ok bool, err error) {

	errResult := models.Returns{}
	errResult.Return.Status = "unauthorized"
	errResult.Return.Message = []string{"access denied"}

	auth, err := w.ApplicationService.AuthenticationService.GetAuthentication(models.Authentication{
		UserName: username,
		Password: password,
	})

	if err != nil {
		errResult.Return.Message = append(errResult.Return.Message, err.Error())
		return false, c.JSON(http.StatusUnauthorized, errResult)
	}

	if auth.Enabled {
		c.Set(ctxKeyAuthUsed, auth)
		return true, nil
	}

	return false, c.JSON(http.StatusUnauthorized, errResult)
}

func (w *WebServer) Logger() echo.MiddlewareFunc {
	return middleware.BodyDump(w.getLogRequestResponse())
}

func (w *WebServer) getLogRequestResponse() middleware.BodyDumpHandler {

	funcBodydump := func(c echo.Context, reqBody, resBody []byte) { go w.buildLog(c, reqBody, resBody) }

	return funcBodydump
}

func (w *WebServer) errorServer(err error, ectx echo.Context) {
	var (
		code      = http.StatusInternalServerError
		errResult models.Returns
	)

	errResult.Return.Status = common.StatusError

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
		log.File(time.Now().Format("errors/2006/01/02/15h.log"), err.Error())
	}
}

func (w *WebServer) buildLog(ectx echo.Context, reqBody, resBody []byte) {

	w.ApplicationService.Utils.BuildLogRequestsReceived(models.ParamsBuildLogsRequestsReceived{
		PathLogs: "requests/2006/01/02/15h.log",
		Ectx:     ectx,
		ReqBody:  reqBody,
		ResBody:  resBody,
	})
}
