package carrier

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/JacobTripp/spiderboss/go/pkg/boundBox"
	"github.com/JacobTripp/spiderboss/go/pkg/movement"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/JacobTripp/spiderboss/go/pkg/winch"
)

type Carrier struct {
	Winches  []*winch.Winch
	BoundBox *boundBox.BoundBox
	Serial   io.ReadWriter
	Logger   *log.Logger
}

func (c *Carrier) ExtendMotorCable(sel int, dis units.Basis) {
}

func (c *Carrier) LineTo(loc units.LocVec, speed float64) {
	c.Logger.Printf("Starting line to %v\n", loc)
	cableLens, err := c.BoundBox.CableLenAt(loc)
	if err != nil {
		c.Logger.Fatal(err)
	}
	movements := make([]movement.Movement, 4)
	for i, cableLen := range cableLens {
		c.Logger.Printf("Cable length for line is %d", cableLen)
		revs, err := c.Winches[i].Spool.Revolutions(cableLen)
		if err != nil {
			c.Logger.Fatal(err)
		}
		if c.Winches[i].CableLength() > cableLen {
			c.Logger.Printf(
				"retracting cable, current length is %d target length is %d",
				c.Winches[i].CableLength(),
				cableLen,
			)
			movements[i].Direction = movement.Backward
		} else {
			c.Logger.Printf(
				"extending cable, current length is %d target length is %d",
				c.Winches[i].CableLength(),
				cableLen,
			)
			movements[i].Direction = movement.Forward
		}
		movements[i].Steps = c.Winches[i].Motor.Steps(revs)
		movements[i].Delay = time.Duration(c.Winches[i].Motor.FastestDuration(revs).Microseconds()/movements[i].Steps) * time.Microsecond

		c.Logger.Printf("spinning motor %d %f times", i, revs)
		_, err = c.Winches[i].SetLength(cableLens[i], 10*time.Second)
		if err != nil {
			c.Logger.Fatal(err)
		}
	}
	normMovements := normalizeTime(movements)
	c.Serial.Write(toBytesMovements(normMovements))
}

func toBytesMovements(mvs []movement.Movement) []byte {
	rt := bytes.NewBufferString("")
	for _, mv := range mvs {
		rt.Write(mv.Bytes())
	}
	rt.WriteString("\n")
	return rt.Bytes()
}

func normalizeTime(mvs []movement.Movement) []movement.Movement {
	longest := movement.Movement{}
	for _, mv := range mvs {
		if mv.Steps > longest.Steps {
			longest = mv
		}
	}
	timeToFinish := longest.Steps * longest.Delay.Microseconds()
	rt := make([]movement.Movement, 4)
	for i, mv := range mvs {
		m := movement.Movement{
			Steps:     mv.Steps,
			Direction: mv.Direction,
			Delay:     time.Duration(timeToFinish/mv.Steps) * time.Microsecond,
		}
		rt[i] = m
	}
	return rt
}

func (c *Carrier) ToHome(speed float64) {
	c.LineTo(units.LocVec{X: 0, Y: 0, Z: 0}, speed)
}

func (c *Carrier) WindCableOnWinch(
	w winch.Winch,
	amount units.Basis,
	speed float64,
) {
}
