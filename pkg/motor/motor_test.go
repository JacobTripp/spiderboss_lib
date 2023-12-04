package motor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
func TestMotor_Forward(t *testing.T) {
	m := &Motor{
		StepsPerRev:  200,
		MinStepDelay: 20 * time.Microsecond,
	}

	move, err := m.Forward(1, 1*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, int64(200), move.Steps)
	assert.Equal(t, int64(5_000), move.Delay.Microseconds())
}
*/

func TestMotor_FastestDuration(t *testing.T) {
	m := &Motor{
		StepsPerRev:  200,
		MinStepDelay: 20 * time.Microsecond,
	}

	dur := m.FastestDuration(1)
	assert.Equal(t, 4000*time.Microsecond, dur)

	dur = m.FastestDuration(10)
	assert.Equal(t, 40000*time.Microsecond, dur)

	dur = m.FastestDuration(1.5)
	assert.Equal(t, 6000*time.Microsecond, dur)
}

func TestMotor_Steps(t *testing.T) {
	m := &Motor{
		StepsPerRev:  200,
		MinStepDelay: 20 * time.Microsecond,
	}
	assert.Equal(t, int64(200), m.Steps(1))
}

/*
type Expected struct {
	value *Movement
	err   error
}

func TestMotor_Move(t *testing.T) {
	testCases := map[string]struct {
		input    float64
		expected Expected
	}{
		"zero": {
			input: 0.0,
			expected: Expected{
				value: &Movement{
					Steps:     0,
					Delay:     20 * time.Microsecond,
					Direction: Forward,
				},
				err: nil,
			},
		},
		"negative one": {
			input: -1.0,
			expected: Expected{
				value: &Movement{
					Steps:     200,
					Delay:     1 * time.Second / 200,
					Direction: Backward,
				},
				err: nil,
			},
		},
		"positive one": {
			input: 1.0,
			expected: Expected{
				value: &Movement{
					Steps:     200,
					Delay:     1 * time.Second / 200,
					Direction: Forward,
				},
				err: nil,
			},
		},
		"too fast": {
			input: 251.0,
			expected: Expected{
				value: &Movement{
					Steps:     200,
					Delay:     1 * time.Second / (200 * 251),
					Direction: Forward,
				},
				err: MovementError,
			},
		},
	}

	m := &Motor{
		StepsPerRev:  200,
		MinStepDelay: 20 * time.Microsecond,
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			movement, err := motor.Move(testCase.input, 1*time.Second)
			if testCase.expected.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected.value, movement)
			} else {
				assert.ErrorIs(t, err, testCase.expected.err)
			}
		})
	}
}
*/
