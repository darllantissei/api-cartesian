package coordinate

import (
	"testing"

	"github.com/darllantissei/api-cartesian/application/models"
	mock_coordinate "github.com/darllantissei/api-cartesian/mocks/coordinate"
	mock_utils "github.com/darllantissei/api-cartesian/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProccess(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUtils := mock_utils.NewMockIUtilsService(ctrl)

	mockPersistence := mock_coordinate.NewMockICoordinatePersistenceCache(ctrl)

	mockUtils.EXPECT().FileExists("").Return(true).AnyTimes()

	mockPersistence.EXPECT().ListPoints().Return(models.Points{{X: 20, Y: 10}}, nil).AnyTimes()

	coordinateService := CoordinateService{
		PersisenceFile: mockPersistence,
		Utils:           mockUtils,
	}

	way, err := coordinateService.Proccess(10, 20)

	assert.Nil(t, err)

	assert.Equal(t, 1, len(way))

	for _, options := range way {

		assert.Equal(t, int(20), options.Distante)

	}
}

func TestCalculateDisctance(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUtils := mock_utils.NewMockIUtilsService(ctrl)

	mockPersistence := mock_coordinate.NewMockICoordinatePersistenceCache(ctrl)

	coordinateService := CoordinateService{
		PersisenceFile: mockPersistence,
		Utils:           mockUtils,
	}

	way := coordinateService.calculateDisctance(10, 20, models.Points{{X: 20, Y: 10}})

	assert.Equal(t, 1, len(way))

	for _, options := range way {

		assert.Equal(t, int(20), options.Distante)

	}
}
