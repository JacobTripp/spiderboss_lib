package spool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidConfig(t *testing.T) {
	testCases := map[string]struct {
		input    *Config
		expected error
	}{
		"valid config": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         2,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        1,
				InitCableOnSpool: 0,
			},
			expected: nil,
		},
		"no empty diameter": {
			input: &Config{
				FullDiam:         2,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        1,
				InitCableOnSpool: 0,
			},
			expected: EmptyDiamError,
		},
		"full diam is same size as empty diam": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         1,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        1,
				InitCableOnSpool: 0,
			},
			expected: FullDiamError,
		},
		"full diam is zero": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         0,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        1,
				InitCableOnSpool: 0,
			},
			expected: FullDiamError,
		},
		"full length is less than empty length": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         2,
				EmptyLength:      1,
				FullLength:       0,
				CableDiam:        1,
				InitCableOnSpool: 0,
			},
			expected: FullLengthError,
		},
		"Cable diam is zero": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         2,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        0,
				InitCableOnSpool: 0,
			},
			expected: CableDiamError,
		},
		"Initial cable on spool is negative": {
			input: &Config{
				EmptyDiam:        1,
				FullDiam:         2,
				EmptyLength:      1,
				FullLength:       1,
				CableDiam:        1,
				InitCableOnSpool: -1,
			},
			expected: InitCableOnSpoolError,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateConfig(testCase.input)
			if testCase.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, testCase.expected)
			}
		})
	}
}
