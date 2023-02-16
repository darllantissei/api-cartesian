package coordinate

import (
	"math"
	"sort"

	"github.com/darllantissei/api-cartesian/application/models"
	statusapplication "github.com/darllantissei/api-cartesian/application/status_application"
)

func (c *CoordinateService) buildError(status statusapplication.StatusApp, messages []string) error {

	retErr := models.Returns{}

	retErr.Return.Message = messages
	retErr.Return.Status = status

	return &retErr
}

func (c *CoordinateService) calculateDistance(coordX, coordY, paramDistance int64, coordinateBase models.Points) []models.Way {

	way := []models.Way{}

	for _, coordinate := range coordinateBase {

		coordParam := []int64{
			coordX,
			coordY,
		}

		coordBase := []int64{
			coordinate.X,
			coordinate.Y,
		}

		distanceXY := c.manhattanDistance(coordParam, coordBase)

		if distanceXY <= paramDistance {

			options := models.Way{
				From: models.Coordinate{
					X: coordX,
					Y: coordY,
				},
				To: models.Coordinate{
					X: coordinate.X,
					Y: coordinate.Y,
				},
				Distante: distanceXY,
			}

			way = append(way, options)
		}

	}

	return way
}

func (c *CoordinateService) sortDistance(way []models.Way) []models.Way {

	sort.Slice(way, func(i, j int) bool {
		return way[i].Distante < way[j].Distante
	})

	return way
}

func (c *CoordinateService) manhattanDistance(vCoordParam, vCoordBase []int64) int64 {

	distance := int64(0)

	for idx, value := range vCoordParam {
		distance += int64(math.Abs(float64(value) - float64(vCoordBase[idx])))
	}

	return distance
}
