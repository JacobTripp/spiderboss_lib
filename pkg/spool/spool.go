package spool

import (
	"errors"
	"fmt"
	"math"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

// Spool is the main type a system should interact with.
// It keeps track of the current cable on the spool and
// does checks when add/removing cable from the spool.
type Spool struct {
	config       *Config
	idealSpool   *IdealSpool
	cableOnSpool units.Basis
}

// Constructor for Spool, this will validate a configuration
// returns an error if the config is bad.
func New(c *Config) (*Spool, error) {
	err := validateConfig(c)
	if err != nil {
		return nil, err
	}

	return &Spool{
		config:       c,
		cableOnSpool: c.InitCableOnSpool, // the validation above guarantees this is set
		idealSpool:   newIdealSpool(c),
	}, nil
}

// Revolutions does two things, it has a side effect of calling AddCable or
// RemoveCable and then it returns the number of revolutions to add/remove
// that amount of cable.
//
// If there isn't enough cable on the spool or if there is too much on the spool
// it will return and error and 0.0 for the revolutions.
func (s *Spool) Revolutions(d units.Basis) (float64, error) {
	if d == 0 {
		return 0.0, nil
	}
	add := d > 0                                // this is used to select if we add or remove cable
	amount := units.Basis(math.Abs(float64(d))) // the sign is captured in the add var

	// Find the first layer than can handle the amount plus the cable
	// already on the spool.
	layer := 0
	cableLeft := amount + s.CableOnSpool()
	for i := 0; i < int(s.idealSpool.numberOfLayers); i++ {
		if s.idealSpool.layers[i].capacity > cableLeft {
			break // this keeps the layer at the previous full layer
		}
		cableLeft -= s.idealSpool.layers[i].capacity
		layer = i
	}

	// instantiate in case the first layer can hold the amount
	leftover := amount
	rowTurns := 0.0

	// if it's not the first layer then set the falues
	if layer > 0 {
		leftover = amount - s.idealSpool.capacityToLayer(layer) // this subraacts all previous layers' capacity
		rowTurns = s.idealSpool.turnsToLayer(layer)             // this give all the revolutions to fill all previous layers
	}
	remainder := float64(leftover) / float64(s.idealSpool.layers[layer].capacityPerRev) // the remaining number of revolutions
	total := rowTurns + remainder

	// if the cable can't be added/removed to the spool then don't do it
	if add {
		err := s.AddCable(amount)
		if err != nil {
			return 0.0, err
		}
	} else {
		err := s.RemoveCable(amount)
		if err != nil {
			return 0.0, err
		}
	}

	return total, nil
}

// If there is less cable on the spool then requested amount to remove
var SpoolEmptyError = errors.New("Cannot remove more cable")

// We don't want to accept negative numbers for the specific add/remove methods
var NegativeCableError = errors.New("Cannot add negative distance to spool")

// If there is too much cable to be added to the spool and the IdealSpool's
// full diameter becomes larger than configured you'll get this error
var SpoolFullError = errors.New("Cannot add more cable to spool")

// This will remove the d amount of cable from the spool. It has the side-effect
// of changing an unexported field that keeps track of the cable on the spool
//
// This checks to make sure there is enough cable to remove before the side-effect
// happens.
func (s *Spool) RemoveCable(d units.Basis) error {
	if d == 0 {
		return nil
	}

	if d < 0 {
		return fmt.Errorf(
			"%w: Please only use positive numbers",
			NegativeCableError,
		)
	}

	if (s.CableOnSpool() - d) < 0 {
		return fmt.Errorf(
			"%w: Attempting to remove too much cable '%d' from spool '%d'",
			SpoolEmptyError,
			d,
			s.CableOnSpool(),
		)
	}

	s.cableOnSpool -= d
	return nil
}

// This will add the d amount of cable to the spool. It has a side-effect of
// changing an unexported field that is used to track the amount of cable on
// the spool.
//
// If the d amount will cause the IdealSpool to overfill it will return an error
// before the side-effect happens
func (s *Spool) AddCable(d units.Basis) error {
	if d == 0 {
		return nil
	}

	if d < 0 {
		return fmt.Errorf(
			"%w: Received Negative distance '%d'",
			NegativeCableError,
			d,
		)
	}

	if d > s.idealSpool.maxCable() {
		return fmt.Errorf(
			"%w: Adding '%d' cable to the existing amount of %d will exceed the limit of %d",
			SpoolFullError,
			d,
			s.cableOnSpool,
			s.idealSpool.maxCable(),
		)
	}

	s.cableOnSpool += d
	return nil
}

// Get the amount of cable on the spool
func (s *Spool) CableOnSpool() units.Basis {
	return s.cableOnSpool
}
