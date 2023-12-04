package spool

import (
	"math"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

// IdealSpool is the assumed best case for winding, no space wasted, and the
// spool fills up uniformly.
type IdealSpool struct {
	numberOfLayers int64
	layers         []*layer
}

type layer struct {
	layerId        int64
	diam           units.Basis
	capacity       units.Basis
	capacityPerRev units.Basis
	maxTurns       float64
}

func (s IdealSpool) capacityToLayer(layer int) units.Basis {
	total := units.Basis(0)
	for i := 0; i <= layer; i++ {
		total += s.layers[i].capacity
	}
	return total
}

func (s IdealSpool) turnsToLayer(layer int) float64 {
	total := 0.0
	for i := 0; i <= layer; i++ {
		total += s.layers[i].maxTurns
	}
	return total
}

func newIdealSpool(c *Config) *IdealSpool {
	rt := &IdealSpool{
		numberOfLayers: int64((c.FullDiam - c.EmptyDiam) / c.CableDiam),
	}
	for layerId := int64(0); layerId < rt.numberOfLayers; layerId++ {
		diam := (units.Basis(layerId) * c.CableDiam) + c.EmptyDiam
		capPerRev := float64(diam) * math.Pi
		lLen := layerLen(
			c.FullLength,
			c.FullDiam,
			diam,
			c.EmptyDiam,
			c.EmptyLength,
		)
		totalCap := capPerRev * float64((lLen / c.CableDiam))
		rt.layers = append(rt.layers, &layer{
			layerId:        layerId,
			diam:           diam,
			capacityPerRev: units.Basis(capPerRev),
			capacity:       units.Basis(totalCap),
			maxTurns:       totalCap / capPerRev,
		})
	}
	return rt
}

func layerLen(
	fullLen,
	fullDiam,
	layerDiam,
	emptyDiam,
	emptyLen units.Basis,
) units.Basis {
	df := float64(fullLen)
	rf := float64(fullDiam)
	ld := float64(layerDiam)
	th := float64(fullDiam - emptyDiam)
	tb := (float64(fullLen) - float64(emptyDiam)) / 2.0

	h := rf - ld
	tan := th / tb
	a := h / tan

	rt := df - 2*a
	return units.Basis(rt)
}

func (s IdealSpool) maxCable() units.Basis {
	rt := units.Basis(0)
	for _, layer := range s.layers {
		rt += layer.capacity
	}
	return rt
}
