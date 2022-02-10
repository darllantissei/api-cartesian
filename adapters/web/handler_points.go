package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (w *WebServer) HandlerGetPoints(c echo.Context) error {

	coordinateX, err := strconv.Atoi(c.QueryParam("coordX"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, w.buildError(errors.New("coordinate X is invalid")))
	}

	coordinateY, err := strconv.Atoi(c.QueryParam("coordY"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, w.buildError(errors.New("coordinate Y is invalid")))
	}

	result, err := w.ApplicationService.CoordinanteService.Proccess(int64(coordinateX), int64(coordinateY))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
