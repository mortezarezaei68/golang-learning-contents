package model

type Level struct {
	Order        int
	Capacity     int
	RemainingCap int
	Products     []Product
}
