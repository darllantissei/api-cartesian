package cachepoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/darllantissei/api-cartesian/application/models"
	"github.com/darllantissei/api-cartesian/application/utils"
)

type PointsSource struct {
	Utils        utils.IUtilsService
	FilePoints   string
	pointsLoaded models.Points
}

func NewPointsCache(config PointsSource) *PointsSource {

	return &PointsSource{
		FilePoints: config.FilePoints,
		Utils:      config.Utils,
	}
}

func (db *PointsSource) ListPoints() (models.Points, error) {

	points := db.pointsLoaded

	if len(points) <= 0 {

		if !db.Utils.FileExists(db.FilePoints) {
			return models.Points{}, errors.New("file with points data not found, please contact the suport")
		}

		content, err := ioutil.ReadFile(db.FilePoints)

		if err != nil {
			return models.Points{}, fmt.Errorf("unable read file with points data. Details: %s", err.Error())
		}

		err = json.Unmarshal(content, &points)

		if err != nil {
			return models.Points{}, fmt.Errorf("unable parse content to struct points data. Details: %s", err.Error())
		}

		db.pointsLoaded = points
	}

	return points, nil
}
