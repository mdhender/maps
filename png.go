package main

import (
	"github.com/fogleman/gg"
	"math"
	"math/rand"
	//"github.com/golang/freetype/truetype"
	//"golang.org/x/image/font/gofont/goregular"
)

func AsPNG(name string, fill bool) error {
	dc := gg.NewContext(width, height)
	dc.SetRGBA(1, 1, 1, 1)
	dc.Clear()

	r := 0.333 // rand.Float64()
	g := 0.333 // rand.Float64()
	b := 0.333 // rand.Float64()
	a := 1.0   // rand.Float64()*0.5 + 0.5
	w := 0.5
	dc.SetLineWidth(w)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y, z := calcz(i+1, j)
			ax, ay := project(x, y, z)
			x, y, z = calcz(i, j)
			bx, by := project(x, y, z)
			x, y, z = calcz(i, j+1)
			cx, cy := project(x, y, z)
			x, y, z = calcz(i+1, j+1)
			dx, dy := project(x, y, z)

			if nan(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			dc.SetRGBA(r, g, b, a)
			dc.MoveTo(ax, ay)
			dc.LineTo(bx, by)
			dc.LineTo(cx, cy)
			dc.LineTo(dx, dy)

			if fill {
				dc.StrokePreserve()
				dc.SetRGBA(1, 1, 1, 1)
				dc.Fill()
			} else {
				dc.Stroke()
			}
		}
	}

	return dc.SavePNG(name)
}

func AsPNG2(name string) error {
	const W = 1024
	const H = 1024
	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	return dc.SavePNG(name)
}

func AsPNG3(name string, fill bool) error {
	dc := gg.NewContext(width, height)
	dc.SetRGBA(1, 1, 1, 1)
	dc.Clear()

	r := 0.333 // rand.Float64()
	g := 0.333 // rand.Float64()
	b := 0.333 // rand.Float64()
	a := 1.0   // rand.Float64()*0.5 + 0.5
	w := 0.5
	dc.SetLineWidth(w)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y, z := calcz3(i+1, j)
			ax, ay := project(x, y, z)
			x, y, z = calcz3(i, j)
			bx, by := project(x, y, z)
			x, y, z = calcz3(i, j+1)
			cx, cy := project(x, y, z)
			x, y, z = calcz3(i+1, j+1)
			dx, dy := project(x, y, z)

			if nan(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			dc.SetRGBA(r, g, b, a)
			dc.MoveTo(ax, ay)
			dc.LineTo(bx, by)
			dc.LineTo(cx, cy)
			dc.LineTo(dx, dy)

			if fill {
				dc.StrokePreserve()
				dc.SetRGBA(1, 1, 1, 1)
				dc.Fill()
			} else {
				dc.Stroke()
			}
		}
	}

	return dc.SavePNG(name)
}

func NewContext(width, height int) *gg.Context {
	dc := gg.NewContext(width, height)
	dc.SetRGBA(1, 1, 1, 1)
	dc.Clear()
	dc.SetLineWidth(0.5)
	return dc
}

func nan(fs ...float64) bool {
	for _, f := range fs {
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return true
		}
	}
	return false
}
