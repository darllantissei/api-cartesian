package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (w *WebServer) HandlerGetPoints(c echo.Context) error {

	var errCollector []string

	coordinateX, err := strconv.Atoi(c.QueryParam("coordX"))

	if err != nil {
		errCollector = append(errCollector, "coordinate X is invalid")
	}

	coordinateY, err := strconv.Atoi(c.QueryParam("coordY"))

	if err != nil {
		errCollector = append(errCollector, "coordinate Y is invalid")
	}

	distance, err := strconv.Atoi(c.QueryParam("distance"))

	if err != nil {
		errCollector = append(errCollector, "the distance is invalid")
	}

	if len(errCollector) > 0 {
		return c.JSON(http.StatusBadRequest, w.buildError(errCollector))
	}

	result, err := w.ApplicationService.CoordinanteService.Proccess(int64(coordinateX), int64(coordinateY), int64(distance))

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if len(result) <= 0 {
		return c.JSON(http.StatusNoContent, nil)
	}

	return c.JSON(http.StatusOK, result)
}
