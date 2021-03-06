package coordinate

import (
	"github.com/darllantissei/api-cartesian/application/common"
	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/darllantissei/api-cartesian/application/utils"
)

type CoordinateService struct {
	PersisenceFile ICoordinatePersistenceFile
	Utils          utils.IUtilsService
}

func (c *CoordinateService) Proccess(coordX, coordY, distance int64) ([]models.Way, error) {

	coordinateBase, err := c.PersisenceFile.ListPoints()

	if err != nil {
		return []models.Way{}, c.buildError(common.StatusError, []string{err.Error()})
	}

	way := c.calculateDistance(coordX, coordY, distance, coordinateBase)

	way = c.sortDistance(way)

	return way, nil
}
