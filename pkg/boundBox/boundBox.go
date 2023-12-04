package boundBox

import (
	"errors"
	"fmt"

	"github.com/JacobTripp/spiderboss/go/pkg/units"
)

type BoundBox struct {
	Origins []units.LocVec
}

func (b *BoundBox) CableLenHome() []units.Basis {
	rt := make([]units.Basis, 4)
	for i, locvec := range b.Origins {
		rt[i] = locvec.Len()
	}
	return rt
}

var BoundBoxError = errors.New("error with bound box")

func (b *BoundBox) CableLenAt(loc units.LocVec) ([]units.Basis, error) {
	rt := make([]units.Basis, 4)
	for i, vec := range b.Origins {
		if b.exceedsBounds(loc) {
			return nil, fmt.Errorf(
				"%w: Location Vector outside of Bound Box",
				BoundBoxError,
			)
		}
		rt[i] = vec.AbsSubtract(loc).Len()
	}
	return rt, nil
}

func (b *BoundBox) exceedsBounds(loc units.LocVec) bool {
	return loc.GreaterThan(b.Origins[2])
}
