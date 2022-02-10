package coordinate

import (
	"errors"
	"math"

	"github.com/darllantissei/api-cartesian/application/models"
)

type CoordinateService struct {
}

func (c *CoordinateService) Proccess(coordiante []int64) ([]models.Coordinate, error) {

	var distance float64 = 0

	coordinateBase := []int64{
		20,
		10,
	}

	for i, value := range coordiante {
		distance += math.Abs(float64(value) - float64(coordinateBase[i]))
	}

	return []models.Coordinate{}, errors.New("yet implementing")
}
