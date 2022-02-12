package application

import (
	"github.com/darllantissei/api-cartesian/application/coordinate"
	"github.com/darllantissei/api-cartesian/application/utils"
)

type Application struct {
	CoordinanteService coordinate.ICoordinateService
	Utils              utils.IUtilsService
}
