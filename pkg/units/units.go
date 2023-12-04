package units

import (
	"math"
)

type Basis int64

// Mathematically speaking, I'm setting the scalar for each unit of measurement.
const (
	Millimeter Basis = 1
	Centimeter Basis = 10
	Meter      Basis = 1000
	Inch       Basis = 25 // this isn't absolutely correct but I don't want floats
	Foot       Basis = 305
	Yard       Basis = 914
)

type LocVec struct {
	X Basis
	Y Basis
	Z Basis
}

func (lv *LocVec) Len() Basis {
	X := float64(lv.X)
	Y := float64(lv.Y)
	Z := float64(lv.Z)
	squares := math.Pow(X, 2) + math.Pow(Y, 2) + math.Pow(Z, 2)
	root := math.Sqrt(squares)
	return Basis(root)
}

func (lv *LocVec) AbsSubtract(loc LocVec) *LocVec {
	return &LocVec{
		X: Basis(math.Abs(float64(lv.X - loc.X))),
		Y: Basis(math.Abs(float64(lv.Y - loc.Y))),
		Z: Basis(math.Abs(float64(lv.Z - loc.Z))),
	}
}

func (lv *LocVec) GreaterThan(loc LocVec) bool {
	return lv.X > loc.X || lv.Y > loc.Y || lv.Z > loc.Z
}
