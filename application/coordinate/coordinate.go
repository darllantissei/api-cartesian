package coordinate

import (
	"math"
	"sort"

	"github.com/darllantissei/api-cartesian/application/models"
)

func (c *CoordinateService) buildError(status string, messages []string) error {

	retErr := models.Returns{}

	retErr.Return.Message = messages
	retErr.Return.Status = status

	return &retErr
}

func (c *CoordinateService) calculateDisctance(coordX, coordY int64, coordinateBase models.Points) []models.Way {

	way := []models.Way{}

	for _, coordinate := range coordinateBase {
		distance := math.Abs(float64(coordinate.X) - float64(coordX))
		distance += math.Abs(float64(coordinate.Y) - float64(coordY))

		options := models.Way{
			From: models.Coordinate{
				X: int(coordX),
				Y: int(coordY),
			},
			To: models.Coordinate{
				X: coordinate.X,
				Y: coordinate.Y,
			},
			Distante: int(distance),
		}

		way = append(way, options)

	}

	return way
}

func (c *CoordinateService) sortDistance(way []models.Way) []models.Way {

	sort.Slice(way, func(i, j int) bool {
		return way[i].Distante < way[j].Distante
	})

	return way
}
