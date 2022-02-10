package coordinate

import "github.com/darllantissei/api-cartesian/application/models"

type ICoordinateService interface {
	Proccess(coordiante []int64) ([]models.Coordinate, error)
}