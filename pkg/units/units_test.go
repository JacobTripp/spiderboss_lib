package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocVec_Len(t *testing.T) {
	testCases := map[string]struct {
		input    LocVec
		expected Basis
	}{
		"[0,0,1] meter": {
			input: LocVec{
				X: 0,
				Y: 0,
				Z: 1 * Meter,
			},
			expected: 1000,
		},
		"[3,0,4] meter": {
			input: LocVec{
				X: 3 * Meter,
				Y: 0,
				Z: 4 * Meter,
			},
			expected: 5000,
		},
		"[3,4,5] meter": {
			input: LocVec{
				X: 3 * Meter,
				Y: 4 * Meter,
				Z: 5 * Meter,
			},
			expected: 7071,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input.Len())
		})
	}
}

func TestLocVec_AbsSubtract(t *testing.T) {
	type Input struct {
		from LocVec
		take LocVec
	}
	testCases := map[string]struct {
		input    Input
		expected *LocVec
	}{
		"10-5": {
			input: Input{
				from: LocVec{
					X: 10 * Meter,
					Y: 10 * Meter,
					Z: 10 * Meter,
				},
				take: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
			},
			expected: &LocVec{
				X: 5 * Meter,
				Y: 5 * Meter,
				Z: 5000,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.expected,
				testCase.input.from.AbsSubtract(testCase.input.take),
			)
		})
	}
}

func TestLocVec_GreaterThan(t *testing.T) {
	type Input struct {
		a LocVec
		b LocVec
	}
	testCases := map[string]struct {
		input    Input
		expected bool
	}{
		"10>5": {
			input: Input{
				a: LocVec{
					X: 10 * Meter,
					Y: 10 * Meter,
					Z: 10 * Meter,
				},
				b: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
			},
			expected: true,
		},
		"5>10": {
			input: Input{
				b: LocVec{
					X: 10 * Meter,
					Y: 10 * Meter,
					Z: 10 * Meter,
				},
				a: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
			},
			expected: false,
		},
		"5>5": {
			input: Input{
				b: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
				a: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
			},
			expected: false,
		},
		"a>b": {
			input: Input{
				a: LocVec{
					X: 10 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
				b: LocVec{
					X: 5 * Meter,
					Y: 5 * Meter,
					Z: 5 * Meter,
				},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.expected,
				testCase.input.a.GreaterThan(testCase.input.b),
			)
		})
	}
}
