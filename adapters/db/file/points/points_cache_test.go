package cachepoints

import (
	"testing"

	mock_utils "github.com/darllantissei/api-cartesian/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListPoints(t *testing.T) {

	pathPoints := "../../../../points.json"

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUtils := mock_utils.NewMockIUtilsService(ctrl)

	mockUtils.EXPECT().FileExists(pathPoints).Return(true).AnyTimes()

	dbCache := NewPointsCache(PointsSource{
		FilePoints: pathPoints,
		Utils:      mockUtils,
	})

	points, err := dbCache.ListPoints()

	assert.Nil(t, err)

	assert.NotEqual(t, 0, len(points))

	for _, coordinate := range points {
		assert.NotEmpty(t, coordinate)
	}
}
