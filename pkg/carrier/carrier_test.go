package carrier

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/JacobTripp/spiderboss/go/pkg/boundBox"
	"github.com/JacobTripp/spiderboss/go/pkg/motor"
	"github.com/JacobTripp/spiderboss/go/pkg/spool"
	"github.com/JacobTripp/spiderboss/go/pkg/units"
	"github.com/JacobTripp/spiderboss/go/pkg/winch"
	"github.com/stretchr/testify/assert"
)

func TestCarrier_LineTo(t *testing.T) {
	tmpFile, err := os.CreateTemp(".", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	logger := log.New(os.Stdout, "", log.Ltime)

	origins := []int{5000, 8542, 9159, 6253}
	winches := make([]*winch.Winch, 4)
	for i, _ := range winches {
		cfg := &spool.Config{
			EmptyDiam:        20 * units.Millimeter,
			FullDiam:         60 * units.Millimeter,
			EmptyLength:      30 * units.Millimeter,
			FullLength:       55 * units.Millimeter,
			CableDiam:        1 * units.Millimeter,
			InitCableOnSpool: 15 * units.Meter,
		}
		spl, err := spool.New(cfg)
		assert.NoError(t, err)
		winches[i] = &winch.Winch{
			Motor: &motor.Motor{
				StepsPerRev:  800,
				MinStepDelay: 100 * time.Microsecond,
			},
			Spool:  spl,
			Origin: units.Basis(origins[i]) * units.Millimeter,
		}
	}

	carrier := Carrier{
		Winches: winches,
		BoundBox: &boundBox.BoundBox{
			Origins: []units.LocVec{
				{
					X: 0,
					Y: 0,
					Z: 250 * units.Centimeter,
				},
				{
					X: 0,
					Y: 550 * units.Centimeter,
					Z: 250 * units.Centimeter,
				},
				{
					X: 280 * units.Centimeter,
					Y: 550 * units.Centimeter,
					Z: 250 * units.Centimeter,
				},
				{
					X: 280 * units.Centimeter,
					Y: 0,
					Z: 250 * units.Centimeter,
				},
			},
		},
		Serial: tmpFile,
		Logger: logger,
	}

	t.Logf("carrier motor delay is %d", carrier.Winches[0].Motor.FastestDuration(1).Microseconds())
	carrier.ToHome(1)
	assert.IsType(t, Carrier{}, carrier)
	carrier.LineTo(units.LocVec{
		X: 100 * units.Centimeter,
		Y: 130 * units.Centimeter,
		Z: 30 * units.Centimeter,
	}, 1)
	carrier.LineTo(units.LocVec{
		X: 100 * units.Centimeter,
		Y: 130 * units.Centimeter,
		Z: 50 * units.Centimeter,
	}, 1)
	fBytes, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.Contains(t, string(fBytes), "{false,43595,203}{false,83497,106}{false,88748,100}{false,59991,147}")
}
