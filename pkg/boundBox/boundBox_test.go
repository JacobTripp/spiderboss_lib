package boundBox

import (
	"testing"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/stretchr/testify/assert"
)

func TestBoundBox_CableLenHome(t *testing.T) {
	testCases := map[string]struct {
		input    *BoundBox
		expected []units.Basis
	}{
		"5x10x2 meter box": {input: &BoundBox{
			Origins: []units.LocVec{
				{
					X: 0 * units.Meter,
					Y: 0 * units.Meter,
					Z: 5 * units.Meter,
				},
				{
					X: 0 * units.Meter,
					Y: 5 * units.Meter,
					Z: 5 * units.Meter,
				},
				{
					X: 5 * units.Meter,
					Y: 5 * units.Meter,
					Z: 5 * units.Meter,
				},
				{
					X: 5 * units.Meter,
					Y: 0 * units.Meter,
					Z: 5 * units.Meter,
				},
			}},
			expected: []units.Basis{5000, 7071, 8660, 7071},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.CableLenHome())
		})
	}
}

type Expected struct {
	err   error
	value []units.Basis
}

func TestBoundBox_CableLenAt(t *testing.T) {
	testCases := map[string]struct {
		input    units.LocVec
		expected Expected
	}{
		"to Middle": {
			input: units.LocVec{
				X: 5000,
				Y: 5 * units.Meter,
				Z: 5 * units.Meter,
			},
			expected: Expected{
				value: []units.Basis{8660, 8660, 8660, 8660},
				err:   nil,
			},
		},
		"too far X": {
			input: units.LocVec{
				X: 10001,
				Y: 5 * units.Meter,
				Z: 5 * units.Meter,
			},
			expected: Expected{
				value: nil,
				err:   BoundBoxError,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			bb := makeCube(t)
			lens, err := bb.CableLenAt(testCase.input)
			if testCase.expected.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected.value, lens)
			} else {
				assert.ErrorIs(t, err, testCase.expected.err)
			}
		})
	}
}

func makeCube(t *testing.T) *BoundBox {
	bb := &BoundBox{
		Origins: []units.LocVec{
			{
				X: 0 * units.Meter,
				Y: 0 * units.Meter,
				Z: 10 * units.Meter,
			},
			{
				X: 0 * units.Meter,
				Y: 10 * units.Meter,
				Z: 10 * units.Meter,
			},
			{
				X: 10 * units.Meter,
				Y: 10 * units.Meter,
				Z: 10 * units.Meter,
			},
			{
				X: 10 * units.Meter,
				Y: 0 * units.Meter,
				Z: 10 * units.Meter,
			},
		},
	}
	return bb
}
