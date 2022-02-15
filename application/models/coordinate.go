package models

type Points []Coordinate

type Coordinate struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Way struct {
	From     Coordinate `json:"from"`
	To       Coordinate `json:"to"`
	Distante int64      `json:"distance"`
}
