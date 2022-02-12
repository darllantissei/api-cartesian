package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {

	utilsService := UtilsService{}

	exists := utilsService.FileExists("../../points.json")

	assert.Equal(t, true, exists)

}
