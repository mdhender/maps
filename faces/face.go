package faces

import (
	"github.com/fogleman/gg"
	"github.com/mdhender/maps/vectors"
	"math"
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

// project to screen coordinates
func (f Faces) Project(width, height int) (proj Faces) {
	halfWidth := float64(width) * 0.5
	halfHeight := float64(height) * 0.5

	for _, face := range f {
		var projected Face
		for i, v := range face {
			projected[i] = vectors.Vector{
				X: (v.X*float64(width))/(2.0*v.W) + halfWidth,
				Y: (v.Y*float64(height))/(2.0*v.W) + halfHeight,
			}
		}
		proj = append(proj, projected)
	}

	//// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	//xyrange := 30.0                         // axis ranges (-xyrange .. +xyrange)
	//xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	//zscale := float64(height) * 0.4         // pixels per z unit (the 0.4 is arbitrary)
	//angle := math.Pi / 6                    // angle of x, y axes (=30degrees)
	//sin30 := math.Sin(angle)                // sin of 30degrees
	//cos30 := math.Cos(angle)                // cos of 30degrees
	//for _, face := range f {
	//	var projected Face
	//	for i, v := range face {
	//		projected[i] = vectors.Vector{
	//			X: halfWidth + (v.X-v.Y)*cos30*xyscale,
	//			Y: halfHeight + (v.X+v.Y)*sin30*xyscale - v.Z*zscale,
	//		}
	//	}
	//	proj = append(proj, projected)
	//}
	//for _, pp := range proj {
	//	log.Printf(" xyscale %f %+v\n", xyscale, pp)
	//}

	return proj
}

func (f Faces) ToPNG(dc *gg.Context) {
	for _, face := range f {
		dc.MoveTo(face[0].X, face[0].Y)
		dc.LineTo(face[1].X, face[1].Y)
		dc.LineTo(face[2].X, face[2].Y)
		dc.Stroke()
	}
}
