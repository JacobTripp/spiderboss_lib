package winch

import (
	"errors"
	"fmt"
	"time"

	"github.com/JacobTripp/spiderboss/go/pkg/motor"
	"github.com/JacobTripp/spiderboss/go/pkg/movement"
	"github.com/JacobTripp/spiderboss/go/pkg/spool"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

type Winch struct {
	Motor       *motor.Motor
	Spool       *spool.Spool
	Origin      units.Basis
	cableLength units.Basis
}

func (w *Winch) Move(d units.Basis, dur time.Duration) (*movement.Movement, error) {
	if d < 0 {
		return w.Retract(d*-1, dur)
	}
	return w.Extend(d, dur)
}

func (w *Winch) CableLength() units.Basis {
	return w.cableLength
}
func (w *Winch) Extend(d units.Basis, dur time.Duration) (*movement.Movement, error) {
	if d == 0 {
		return new(movement.Movement), nil
	}

	revs, err := w.Spool.Revolutions(d)
	if err != nil {
		return nil, err
	}

	w.cableLength += d
	rt := new(movement.Movement)
	rt.Steps = w.Motor.Steps(revs)
	rt.Delay = w.Motor.FastestDuration(revs)
	rt.Direction = movement.Forward
	return rt, nil
}

func (w *Winch) Retract(d units.Basis, dur time.Duration) (*movement.Movement, error) {
	if d == 0 {
		return new(movement.Movement), nil
	}
	if d < 0 {
		d = d * -1
	}
	if w.cableLength < d {
		return nil, fmt.Errorf(
			"%w: retracting more (%d) than available cable (%d)",
			WinchError,
			d,
			w.cableLength,
		)
	}

	revs, err := w.Spool.Revolutions(d)
	if err != nil {
		return nil, err
	}
	w.cableLength -= d

	rt := new(movement.Movement)
	rt.Steps = w.Motor.Steps(revs)
	rt.Delay = w.Motor.FastestDuration(revs)
	rt.Direction = movement.Backward
	return rt, nil
}

var WinchError = errors.New("Invalid winch command")

func (w *Winch) SetLength(d units.Basis, dur time.Duration) (*movement.Movement, error) {
	if d < 0 {
		return nil, fmt.Errorf(
			"%w: Cannot set length of cable to negative '%d'",
			WinchError,
			d,
		)
	}
	dist := d - w.cableLength
	return w.Move(dist, dur)
}
