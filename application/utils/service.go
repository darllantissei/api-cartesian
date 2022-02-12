package utils

import (
	"os"
)

type UtilsService struct{}

func (u *UtilsService) FileExists(fileName string) bool {

	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()

}
