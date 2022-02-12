package web

import (
	"fmt"

	"github.com/darllantissei/api-cartesian/application"
	"github.com/darllantissei/api-cartesian/application/common"
	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type WebServer struct {
	ApplicationService application.Application
}

func MakeNewWebServer(applicationService application.Application) *WebServer {
	return &WebServer{
		ApplicationService: applicationService,
	}
}

func (w *WebServer) Serve(port int, isDebug bool) {

	e := echo.New()

	e.HideBanner = true

	if isDebug {
		e.Debug = true
		e.Use(middleware.Logger())
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = w.errorServer

	w.buildRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))

}

func (w *WebServer) buildError(errs []string) error {

	if len(errs) > 0 {

		errResult := models.Returns{}

		errResult.Return.Status = common.StatusError
		errResult.Return.Message = errs

		return errResult
	}

	return nil
}
