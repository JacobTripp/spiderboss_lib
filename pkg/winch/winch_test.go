package winch

import (
	"testing"
	"time"

	"github.com/JacobTripp/spiderboss/go/pkg/motor"
	"github.com/JacobTripp/spiderboss/go/pkg/movement"
	"github.com/JacobTripp/spiderboss/go/pkg/spool"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/stretchr/testify/assert"
)

var mockMovement = &movement.Movement{
	Steps:     0,
	Delay:     0,
	Direction: movement.Forward,
}

type Expected struct {
	err      error
	cableLen units.Basis
}

type TestCase struct {
	input    units.Basis
	expected Expected
}

func TestWinch_Move(t *testing.T) {
	testCases := map[string]TestCase{
		"0mm": {
			input: 0 * units.Millimeter,
			expected: Expected{
				cableLen: 150 * units.Centimeter,
				err:      nil,
			},
		},
		"5cm": {
			input: 5 * units.Centimeter,
			expected: Expected{
				cableLen: 155 * units.Centimeter,
				err:      nil,
			},
		},
		"-1m": {
			input: -1 * units.Meter,
			expected: Expected{
				cableLen: 50 * units.Centimeter,
				err:      nil,
			},
		},
		"-1.6m": {
			input: -160 * units.Centimeter,
			expected: Expected{
				cableLen: 150 * units.Centimeter,
				err:      WinchError,
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			spoolConfig := &spool.Config{
				EmptyDiam:        10 * units.Millimeter,
				FullDiam:         20 * units.Millimeter,
				EmptyLength:      10 * units.Millimeter,
				FullLength:       10 * units.Millimeter,
				CableDiam:        1 * units.Millimeter,
				InitCableOnSpool: 1 * units.Meter,
			}
			spool, err := spool.New(spoolConfig)
			assert.NoError(t, err)
			motor := &motor.Motor{
				StepsPerRev:  200,
				MinStepDelay: 50 * time.Microsecond,
			}

			w := &Winch{
				Spool:  spool,
				Motor:  motor,
				Origin: 5 * units.Meter,
			}

			w.Extend(150*units.Centimeter, 1*time.Second)

			_, err = w.Move(testCase.input, 1*time.Second)
			assert.Equal(t, testCase.expected.cableLen, w.cableLength)
			assert.ErrorIs(t, err, testCase.expected.err)
		})
	}
}

func TestWinch_SetLength(t *testing.T) {
	testCases := map[string]TestCase{
		"0mm": {
			input: 0 * units.Millimeter,
			expected: Expected{
				cableLen: 0,
				err:      nil,
			},
		},
		"5cm": {
			input: 5 * units.Centimeter,
			expected: Expected{
				cableLen: 50,
				err:      nil,
			},
		},
		"-1m": {
			input: -1 * units.Meter,
			expected: Expected{
				cableLen: 0,
				err:      WinchError,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			spoolConfig := &spool.Config{
				EmptyDiam:        10 * units.Millimeter,
				FullDiam:         20 * units.Millimeter,
				EmptyLength:      10 * units.Millimeter,
				FullLength:       10 * units.Millimeter,
				CableDiam:        1 * units.Millimeter,
				InitCableOnSpool: 1 * units.Meter,
			}
			spool, err := spool.New(spoolConfig)
			assert.NoError(t, err)
			motor := &motor.Motor{
				StepsPerRev:  200,
				MinStepDelay: 50 * time.Microsecond,
			}

			w := &Winch{
				Spool:  spool,
				Motor:  motor,
				Origin: 5 * units.Meter,
			}

			_, err = w.SetLength(testCase.input, 1*time.Second)
			assert.Equal(t, testCase.expected.cableLen, w.cableLength)
			assert.ErrorIs(t, err, testCase.expected.err)
		})
	}
}
