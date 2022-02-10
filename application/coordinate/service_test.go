package coordinate

import "testing"

func TestProcess(t *testing.T) {

	coordTest := CoordinateService{}

	coordTest.Proccess([]int64{10,20})
}
