package coordinate

import "github.com/darllantissei/api-cartesian/application/models"

type ICoordinateService interface {
	Proccess(coordX, coordY, distance int64) ([]models.Way, error)
}

type ICoordinateReaderFile interface {
	ListPoints() (models.Points, error)
}

type ICoordinatePersistenceFile interface {
	ICoordinateReaderFile
}
