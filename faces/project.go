package faces

import (
	"github.com/mdhender/maps/vectors"
)

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
