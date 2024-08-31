package types

type Position struct {
	X, Y int
}

type Entity struct {
	Name    string
	Id      string
	Version string
	Position
}

type SpecialEntity struct {
	Entity
	SpecialField float64
}