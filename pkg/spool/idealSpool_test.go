package spool

import (
	"testing"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/stretchr/testify/assert"
)

func TestFindMaxCable(t *testing.T) {
	testCases := map[string]struct {
		input    *Config
		expected int
	}{
		"one tiny loop": {
			input: &Config{
				EmptyDiam:   1 * units.Millimeter,
				FullDiam:    2 * units.Millimeter,
				EmptyLength: 1 * units.Millimeter,
				FullLength:  1 * units.Millimeter,
				CableDiam:   units.Millimeter,
			},
			expected: 3,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			s := newIdealSpool(testCase.input)
			assert.Equal(t, testCase.expected, int(s.maxCable()))
		})
	}
}

func TestIdealSpool(t *testing.T) {
	config := &Config{
		EmptyDiam:   10 * units.Millimeter,
		FullDiam:    20 * units.Millimeter,
		EmptyLength: 10 * units.Millimeter,
		FullLength:  20 * units.Millimeter,
		CableDiam:   units.Millimeter,
	}
	spool := newIdealSpool(config)

	assert.Equal(t, 10, int(spool.numberOfLayers))

	// first layer
	assert.Equal(t, 314, int(spool.layers[0].capacity))
	assert.Equal(t, 31, int(spool.layers[0].capacityPerRev))

	// middle layer
	assert.Equal(t, 615, int(spool.layers[4].capacity))
	assert.Equal(t, 43, int(spool.layers[4].capacityPerRev))

	// last layer
	assert.Equal(t, 1134, int(spool.layers[9].capacity))
	assert.Equal(t, 59, int(spool.layers[9].capacityPerRev))
}
