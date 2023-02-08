package main

import (
	"bytes"
	"fmt"
	"github.com/mdhender/maps/draw"
	"log"
	"math"
	"os"
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
	//if err := AsSVG("out.svg"); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created out.svg\n")
	//if err := AsPNG2("out.png"); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created out.png\n")
	//if err := AsPNG("fill.png", true); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created fill.png\n")
	//if err := AsPNG("nofill.png", false); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created nofill.png\n")
	//if err := AsPNG3("nocirc.png", false); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created nocirc.png\n")
	//if err := AsPNG4("vector.png", false, Vectors()); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("created vector.png\n")

	// compose an object from multiple faces
	object := MakeHouse()
	log.Printf("faces: object    %4d\n", len(object))
	//// clip the object
	//clipped := object.Clip(100, 100)
	//log.Printf("faces: clipped   %4d\n", len(clipped))
	//// project the clipped object
	//projected := clipped.Project(100, 100)
	//log.Printf("faces: projected %4d\n", len(projected))
	// create an SVG
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("<svg style='stroke:grey; fill:none; stroke-width:1' width='%d' height='%d' xmlns='http://www.w3.org/2000/svg'>\n", 128, 64))
	object.ToSVG(buf, 64)
	buf.WriteString("</svg>")
	if err := os.WriteFile("faces.svg", buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	log.Printf("created faces.svg\n")

	// save as PNG
	dc := draw.NewContext(100, 100)
	r := 0.333 // rand.Float64()
	g := 0.333 // rand.Float64()
	b := 0.333 // rand.Float64()
	a := 1.0   // rand.Float64()*0.5 + 0.5
	dc.SetRGBA(r, g, b, a)
	object.ToPNG(dc)
	//projected.ToPNG(dc)
	if err := dc.SavePNG("faces.png"); err != nil {
		log.Fatal(err)
	}
	log.Printf("created faces.png\n")
}
