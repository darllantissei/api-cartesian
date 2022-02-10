package web

import (
	"github.com/darllantissei/api-cartesian/application"
	"github.com/darllantissei/api-cartesian/application/common"
	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/eucatur/go-toolbox/api"
)

type WebServer struct {
	ApplicationService application.Application
}

func MakeNewWebServer(applicationService application.Application) *WebServer {
	return &WebServer{
		ApplicationService: applicationService,
	}
}

func (w *WebServer) Serve(port int) {

	echoServer := api.Make()

	api.Use(w.Logger())

	echoServer.HTTPErrorHandler = w.errorServer

	api.ProvideEchoInstance(w.buildRoutes)

	api.Run()
}

func (w *WebServer) buildError(err error) error {

	if err != nil {

		errResult := models.Returns{}

		errResult.Return.Status = common.StatusError
		errResult.Return.Message = []string{err.Error()}

		return errResult
	}

	return nil
}
