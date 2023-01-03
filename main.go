package main

import (
	"log"
	"math"
)

// From The Go Programming Language book

const (
	//width, height = 600, 320            // canvas size in pixels
	//cells         = 100                 // number of grid cells
	width, height = 1200, 640           // canvas size in pixels
	cells         = 200                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange .. +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit (the 0.4 is arbitrary)
	angle         = math.Pi / 6         // angle of x, y axes (=30degrees)
)

var (
	sin30 = math.Sin(angle) // sin of 30degrees
	cos30 = math.Cos(angle) // cos of 30degrees
)

func main() {
	if err := AsSVG("out.svg"); err != nil {
		log.Fatal(err)
	}
	log.Printf("created out.svg\n")
	if err := AsPNG2("out.png"); err != nil {
		log.Fatal(err)
	}
	log.Printf("created out.png\n")
	if err := AsPNG("fill.png", true); err != nil {
		log.Fatal(err)
	}
	log.Printf("created fill.png\n")
	if err := AsPNG("nofill.png", false); err != nil {
		log.Fatal(err)
	}
	log.Printf("created nofill.png\n")
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
