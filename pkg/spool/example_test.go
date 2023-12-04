package spool_test

import (
	"fmt"

	"github.com/JacobTripp/spiderboss/go/pkg/spool"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

func Example() {
	config := &spool.Config{
		EmptyDiam:        20,
		FullDiam:         60,
		EmptyLength:      30,
		FullLength:       50,
		CableDiam:        1,
		InitCableOnSpool: 0,
	}

	s, err := spool.New(config)
	if err != nil {
		panic(err)
	}

	err = s.AddCable(10 * units.Centimeter)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.CableOnSpool()) // First line of Output

	err = s.RemoveCable(10 * units.Centimeter)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.CableOnSpool()) // Second line of Output

	revs, err := s.Revolutions(31 * units.Millimeter)
	if err != nil {
		panic(err)
	}
	fmt.Println(revs)             // Third line of Output
	fmt.Println(s.CableOnSpool()) // Fourth line of Output

	// Output:
	// 100
	// 0
	// 0.5
	// 31
}
