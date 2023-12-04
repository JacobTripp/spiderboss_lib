package movement

import (
	"fmt"
	"time"
)

type Direction int

const (
	Forward  Direction = iota
	Backward           = iota
)

type Movement struct {
	Steps     int64
	Delay     time.Duration
	Direction Direction
}

func (mv *Movement) Bytes() []byte {
	dir := true
	if mv.Direction == Forward {
		dir = false
	}
	return []byte(
		fmt.Sprintf(
			"{%t,%d,%d}",
			dir,
			mv.Steps,
			mv.Delay.Microseconds(),
		),
	)
}

func EmptyMovement() *Movement {
	return new(Movement)
}
