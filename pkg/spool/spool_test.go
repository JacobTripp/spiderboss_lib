package spool

import (
	"fmt"
	"testing"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/stretchr/testify/assert"
)

func TestRevolutions(t *testing.T) {
	type Expected struct {
		revs    float64
		onSpool units.Basis
		err     error
	}
	testCases := map[int]struct {
		input    units.Basis
		expected Expected
	}{
		0: {
			input: 0 * units.Millimeter,
			expected: Expected{
				revs:    0,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		1: {
			input: -0 * units.Millimeter,
			expected: Expected{
				revs:    0,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		2: {
			input: 4_537 * units.Millimeter, // this is max, should be about 4537
			expected: Expected{
				revs:    100,
				onSpool: 4_537 * units.Millimeter,
				err:     nil,
			},
		},
		3: {
			input: -4_537 * units.Millimeter, // this is max, should be about 4537
			expected: Expected{
				revs:    100,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		4: {
			input: 31 * units.Millimeter,
			expected: Expected{
				revs:    1,
				onSpool: 31 * units.Millimeter,
				err:     nil,
			},
		},
		5: {
			input: -31 * units.Millimeter,
			expected: Expected{
				revs:    1,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		6: {
			input: 62 * units.Millimeter,
			expected: Expected{
				revs:    2,
				onSpool: 62 * units.Millimeter,
				err:     nil,
			},
		},
		7: {
			input: -62 * units.Millimeter,
			expected: Expected{
				revs:    2,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		8: {
			input: 15 * units.Millimeter,
			expected: Expected{
				revs:    0.5,
				onSpool: 15 * units.Millimeter,
				err:     nil,
			},
		},
		9: {
			input: -7 * units.Millimeter,
			expected: Expected{
				revs:    0.25,
				onSpool: 8 * units.Millimeter,
				err:     nil,
			},
		},
		10: {
			input: -8 * units.Millimeter,
			expected: Expected{
				revs:    0.27,
				onSpool: 0 * units.Millimeter,
				err:     nil,
			},
		},
		11: {
			input: 659 * units.Millimeter,
			expected: Expected{
				revs:    20,
				onSpool: 659,
				err:     nil,
			},
		},
		12: {
			input: -659 * units.Millimeter,
			expected: Expected{
				revs:    20,
				onSpool: 0,
				err:     nil,
			},
		},
		13: {
			input: 314 * units.Millimeter,
			expected: Expected{
				revs:    10.1,
				onSpool: 314,
				err:     nil,
			},
		},
	}

	config := &Config{
		EmptyDiam:        10 * units.Millimeter,
		FullDiam:         20 * units.Millimeter,
		EmptyLength:      10 * units.Millimeter,
		FullLength:       10 * units.Millimeter,
		CableDiam:        1 * units.Millimeter,
		InitCableOnSpool: 0 * units.Millimeter,
	}
	/*
		FYI: These tests use the same spool
	*/
	spool, err := New(config)
	if !assert.NoError(t, err) {
		t.Fatal("Could not instantiate a spool!")
	}
	// Need to do this to ensure order of tests since maps don't guarantee order
	for i := 0; i < len(testCases); i++ {
		t.Run(fmt.Sprintf("running test %d", i), func(t *testing.T) {
			revs, err := spool.Revolutions(testCases[i].input)
			if err != nil {
				assert.ErrorIs(t, err, testCases[i].expected.err)
			}
			assert.InDelta(t, testCases[i].expected.revs, revs, 0.5)
			assert.Equal(t, testCases[i].expected.onSpool, spool.CableOnSpool())
		})
	}
}

func TestCableOnSpoolRemove(t *testing.T) {
	config := &Config{
		EmptyDiam:        10 * units.Millimeter,
		FullDiam:         20 * units.Millimeter,
		EmptyLength:      30 * units.Millimeter,
		FullLength:       30 * units.Millimeter,
		CableDiam:        units.Millimeter,
		InitCableOnSpool: 1000 * units.Millimeter,
	}
	testCases := map[string]struct {
		input    int64
		expected int64
	}{
		"remove Zero": {
			input:    0,
			expected: 1000,
		},
		"remove one": {
			input:    1,
			expected: 999,
		},
		"remove all": {
			input:    1000,
			expected: 0,
		},
	}
	for name, testCase := range testCases {
		spool, err := New(config)
		if !assert.NoError(t, err) {
			t.Fatal(err)
		}
		_ = spool.RemoveCable(units.Basis(testCase.input) * units.Millimeter)
		t.Run(name, func(t *testing.T) {
			assert.Equal(
				t,
				units.Basis(testCase.expected)*units.Millimeter,
				spool.CableOnSpool(),
			)
		})
	}
}

func TestCableOnSpoolAdd(t *testing.T) {
	config := &Config{
		EmptyDiam:   10 * units.Millimeter,
		FullDiam:    20 * units.Millimeter,
		EmptyLength: 30 * units.Millimeter,
		FullLength:  30 * units.Millimeter,
		CableDiam:   units.Millimeter,
	}
	testCases := map[string]struct {
		input    int64
		expected int64
	}{
		"add Zero": {
			input:    0,
			expected: 0,
		},
		"add one": {
			input:    1,
			expected: 1,
		},
	}
	for name, testCase := range testCases {
		spool, err := New(config)
		if !assert.NoError(t, err) {
			t.Fatal(err)
		}
		spool.AddCable(units.Basis(testCase.input) * units.Millimeter)
		t.Run(name, func(t *testing.T) {
			assert.Equal(
				t,
				units.Basis(testCase.expected)*units.Millimeter,
				spool.CableOnSpool(),
			)
		})
	}
}

func TestRemoveCable(t *testing.T) {
	testCases := map[string]struct {
		input    int64
		expected error
	}{
		"remove one": {
			input:    1,
			expected: nil,
		},
		"remove all": {
			input:    1000,
			expected: nil,
		},
		"remove negative": {
			input:    -1,
			expected: NegativeCableError,
		},
		"remove too much": {
			input:    1001,
			expected: SpoolEmptyError,
		},
		"remove way too much": {
			input:    9999,
			expected: SpoolEmptyError,
		},
	}

	config := &Config{
		EmptyDiam:        10 * units.Millimeter,
		FullDiam:         20 * units.Millimeter,
		EmptyLength:      30 * units.Millimeter,
		FullLength:       30 * units.Millimeter,
		CableDiam:        units.Millimeter,
		InitCableOnSpool: 1000 * units.Millimeter,
	}

	for name, testCase := range testCases {
		spool, err := New(config)
		if !assert.NoError(t, err) {
			t.Fatal(err)
		}
		t.Run(name, func(t *testing.T) {
			assert.ErrorIs(
				t,
				spool.RemoveCable(units.Basis(testCase.input)*units.Millimeter),
				testCase.expected,
			)
		})
	}
}

func TestAddCable(t *testing.T) {
	testCases := map[string]struct {
		input    int64
		expected error
	}{
		"add one": {
			input:    1,
			expected: nil,
		},
		"add negative": {
			input:    -1,
			expected: NegativeCableError,
		},
		"add too much": {
			input:    9897,
			expected: SpoolFullError,
		},
		"add way too much": {
			input:    999999999999999,
			expected: SpoolFullError,
		},
	}

	config := &Config{
		EmptyDiam:   10 * units.Millimeter,
		FullDiam:    20 * units.Millimeter,
		EmptyLength: 30 * units.Millimeter,
		FullLength:  30 * units.Millimeter,
		CableDiam:   units.Millimeter,
	}

	for name, testCase := range testCases {
		spool, err := New(config)
		if !assert.NoError(t, err) {
			t.Fatal(err)
		}
		t.Run(name, func(t *testing.T) {
			assert.ErrorIs(
				t,
				spool.AddCable(units.Basis(testCase.input)*units.Millimeter),
				testCase.expected,
			)
		})
	}
}
