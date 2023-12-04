package main_test

/*
import (
	"time"

	"github.com/JacobTripp/spiderboss/go/pkg/motor"
	"github.com/JacobTripp/spiderboss/go/pkg/spool"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/JacobTripp/spiderboss/go/pkg/winch"
)

func Example() {
	spoolConfig := &spool.Config{
		EmptyDiam:        10 * units.Millimeter,
		FullDiam:         15 * units.Millimeter,
		EmptyLength:      10 * units.Millimeter,
		FullLength:       15 * units.Millimeter,
		CableDiam:        1 * units.Millimeter,
		InitCableOnSpool: 10 * units.Meter,
	}

	motorConfig := &motor.Config{
		StepsPerRev:  200,
		MinStepDelay: 20 * time.Microsecond,
	}

	motor1, _ := motor.New(motorConfig)
	spool1, _ := spool.New(spoolConfig)

	winchConfig := &winch.Config{
		Spool: spool1,
		Motor: motor1,
	}
	winch1, _ := winch.New(winchConfig)

	bbConfig := &boundBox.Config{
		Origin:  []int64{1 * units.Meter},
	}

	bb, _ := boundBox.New(bbConfig)
	serialConn, _ := serial.New("/foo/bar", 12800)

	carrierConfig := &carrier.Config{
		Winches: []*winch.Winch{winch1},
		BoundBox: bb,
		Serial: serialConn,
	}
	car := carrier.New(carrierConfig)

	speed := 0.8 // 80% of max speed
	// These should all take x,y,z coords
	_ = car.LineTo(LocVec{1 * units.Millimeter}, speed)
	_ = car.ArcTo(LocVec{15 * units.Millimeter}, speed) // this doesn't make sense in a 1D set up fyi
	_ = car.Move(carrier.Segment{LocVec{1 * units.Millimeter}, speed, carrier.Line})
	_ = car.Home(speed)
	_ = car.Path([]carrier.Segment{
		{LocVec{1 * units.Meter}, speed, carrier.Line},
		{LocVec{50 * units.Centimeter}, speed, carrier.Arc},
		{LocVec{1 * units.Meter, speed}, carrier.Line},
	})
	_ = car.FastHome() // same as carrier.Home(1)
}

*/
