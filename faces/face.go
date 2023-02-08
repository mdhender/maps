package faces

import (
	"github.com/mdhender/maps/vectors"
)

type Faces []Face

// Face is the three coordinates
type Face [3]vectors.Vector

func New(x1, y1, z1 float64, x2, y2, z2 float64, x3, y3, z3 float64) Face {
	return Face{
		vectors.NewVector(x1, y1, z1),
		vectors.NewVector(x2, y2, z2),
		vectors.NewVector(x3, y3, z3),
	}
}
