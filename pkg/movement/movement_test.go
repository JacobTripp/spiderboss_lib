package movement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMovement_Bytes(t *testing.T) {
	m := Movement{
		Steps:     3000,
		Delay:     100_000 * time.Microsecond,
		Direction: Forward,
	}
	assert.Equal(t, "{false,3000,100000}", string(m.Bytes()))
	m = Movement{
		Steps:     3000,
		Delay:     100_000 * time.Microsecond,
		Direction: Backward,
	}
	assert.Equal(t, "{true,3000,100000}", string(m.Bytes()))
}
