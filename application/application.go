package application

import (
	"github.com/darllantissei/api-cartesian/application/authentication"
	"github.com/darllantissei/api-cartesian/application/coordinate"
	"github.com/darllantissei/api-cartesian/application/utils"
)

type Application struct {
	AuthenticationService authentication.IAuthenticationService
	CoordinanteService    coordinate.ICoordinateService
	Utils                 utils.IUtilsService
}
