package models

type Points []Coordinate

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Way struct {
	From     Coordinate `json:"from"`
	To       Coordinate `json:"to"`
	Distante int        `json:"distance"`
}
