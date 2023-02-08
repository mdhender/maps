package draw

import "math"

var (
	sin30 = math.Sin(angle) // sin of 30degrees
	cos30 = math.Cos(angle) // cos of 30degrees
)

func calcz(i, j int) (x, y, z float64) {
	// find point (x,y) at corner of cell (i,j)
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)
	// compute surface height z
	z = f(x, y)
	// return the coordinates
	return x, y, z
}

func calcz3(i, j int) (x, y, z float64) {
	x, y = scale(i, j)
	// compute surface height z
	z = f(x, y) //math.Sqrt(float64(i*i + j*j))
	// return the coordinates
	return x, y, z
}

func corner(i, j int) (sx, sy float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// f returns the z for a given x and y
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from the origin
	return math.Sin(r) / r
}

// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
func project(x, y, z float64) (sx, sy float64) {
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// find point (x,y) at corner of cell (i,j)
func scale(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return x, y
}
