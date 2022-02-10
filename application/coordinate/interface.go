package coordinate

import "github.com/darllantissei/api-cartesian/application/models"

type ICoordinateService interface {
	Proccess(coordX, coordY int64) ([]models.Coordinate, error)
}