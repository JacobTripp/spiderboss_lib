package spool

import (
	"errors"
	"fmt"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

type Config struct {
	EmptyDiam        units.Basis
	FullDiam         units.Basis
	EmptyLength      units.Basis
	FullLength       units.Basis
	CableDiam        units.Basis
	InitCableOnSpool units.Basis
}

// Various errors with configuration will return one of these
var (
	InvalidSpoolConfigError = errors.New("The config is not valid for a spool")
	EmptyDiamError          = errors.New("invalid empty diameter")
	FullDiamError           = errors.New("invalid full diameter")
	EmptyLengthError        = errors.New("invalid empty length")
	FullLengthError         = errors.New("invalid full length")
	CableDiamError          = errors.New("invalid cable diameter")
	InitCableOnSpoolError   = errors.New("invalid amount of cable on spool")
)

// This will painfully return each error as you fix them, it doesn't report all
// errors in one go.
func validateConfig(c *Config) error {
	if c.EmptyDiam <= 0 {
		return fmt.Errorf(
			"%w: Wanted value > 0, got %d",
			EmptyDiamError,
			c.EmptyDiam,
		)
	}

	if c.FullDiam <= 0 ||
		c.FullDiam <= c.EmptyDiam {
		return fmt.Errorf(
			"%w: Wanted 0 < FullDiam > EmptyDiam, got %d",
			FullDiamError,
			c.FullDiam,
		)
	}

	if c.EmptyLength <= 0 {
		return fmt.Errorf(
			"%w: Wanted value > 0, got %d",
			EmptyLengthError,
			c.EmptyLength,
		)
	}

	if c.FullLength <= 0 ||
		c.FullLength < c.EmptyLength {
		return fmt.Errorf(
			"%w: Wanted value 0 < FullLength >= EmptyLength, got %d",
			FullLengthError,
			c.FullLength,
		)
	}

	if c.CableDiam <= 0 {
		return fmt.Errorf(
			"%w: Wanted value > 0, got %d",
			CableDiamError,
			c.CableDiam,
		)
	}

	if c.InitCableOnSpool < 0 {
		return fmt.Errorf(
			"%w: Wanted value > 0, got %d",
			InitCableOnSpoolError,
			c.InitCableOnSpool,
		)
	}
	return nil
}
