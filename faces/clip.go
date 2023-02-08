package faces

import (
	"github.com/mdhender/maps/vectors"
	"math"
)

func (f Faces) Clip(width, height int) (clip Faces) {
	aspect := float64(width) / float64(height)
	clipMatrix := vectors.SetupClip(80.0*(math.Pi/180.0), aspect)

	for _, face := range f {
		var clipped Face
		for i, v := range face {
			// Don't get confused here. I assume the divide leaves v.w alone.
			clipped[i] = vectors.MatrixMultiply(v, clipMatrix).Div(v.W)
		}
		clip = append(clip, clipped)
	}

	return clip
}
