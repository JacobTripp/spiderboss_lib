package motor

import (
	"math"
	"time"
)

type Motor struct {
	StepsPerRev  int64
	MinStepDelay time.Duration
}

func (m *Motor) FastestDuration(revs float64) time.Duration {

	return time.Duration(m.MinStepDelay.Microseconds()*int64((revs*float64(m.StepsPerRev)))) * time.Microsecond
}

func (m *Motor) Steps(revs float64) int64 {
	return int64(math.Abs(revs * float64(m.StepsPerRev)))
}

/*
var MovementError = errors.New("Bad movement command")

func (m *Motor) Forward(revs float64, byTime time.Duration) (*Movement, error) {
	return m.move(revs, byTime, Forward)
}

func (m *Motor) Backward(revs float64, byTime time.Duration) (*Movement, error) {
	return m.move(revs, byTime, Backward)
}

func (m *Motor) Move(revs float64, byTime time.Duration) (*Movement, error) {
	if revs < 0 {
		return m.Backward(revs, byTime)
	}
	return m.Forward(revs, byTime)
}

func (m *Motor) move(revs float64, byTime time.Duration, dir Direction) (*Movement, error) {
	if revs == 0.0 {
		return &Movement{
			Steps:     0,
			Delay:     m.MinStepDelay,
			Direction: Forward,
		}, nil
	}

	steps := m.Steps(revs)
	delay := time.Duration(byTime.Microseconds()/steps) * time.Microsecond
	if delay < m.MinStepDelay {
		return nil, fmt.Errorf(
			"%w: Moving %f revolutions in %d time is too fast",
			MovementError,
			revs,
			byTime,
		)
	}
	return &Movement{
		Steps:     steps,
		Delay:     time.Duration(byTime.Microseconds()/steps) * time.Microsecond,
		Direction: dir,
	}, nil
}
*/
