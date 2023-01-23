package main

import (
	"github.com/mdhender/maps/vectors"
	"math"
)

const M_PI = math.Pi
const NEAR_Z = 0.5
const FAR_Z = 2.5

func Vectors() []vectors.Vector {
	var v []vectors.Vector
	// front face
	v = append(v, vectors.NewVector(0, 0, 30))
	v = append(v, vectors.NewVector(16, 0, 30))
	v = append(v, vectors.NewVector(16, 10, 30))
	v = append(v, vectors.NewVector(8, 16, 30))
	v = append(v, vectors.NewVector(0, 10, 30))
	v = append(v, vectors.NewVector(0, 0, 30))

	// rear face
	v = append(v, vectors.NewVector(0, 0, 54))
	v = append(v, vectors.NewVector(16, 0, 54))
	v = append(v, vectors.NewVector(16, 10, 54))
	v = append(v, vectors.NewVector(8, 16, 54))
	v = append(v, vectors.NewVector(0, 10, 54))
	v = append(v, vectors.NewVector(0, 0, 54))

	//// beams
	//v = append(v, NewVector(0, 0, 30))
	//v = append(v, NewVector(0, 0, 54))
	//v = append(v, NewVector(16, 0, 30))
	//v = append(v, NewVector(16, 0, 54))

	return ProjectAndClip(100, 100, v)
}

func ProjectAndClip(width, height int, vertex []vectors.Vector) (dst []vectors.Vector) {
	halfWidth := float64(width) * 0.5
	halfHeight := float64(height) * 0.5
	aspect := float64(width) / float64(height)
	clipMatrix := SetupClipMatrix(60.0*(math.Pi/180.0), aspect)

	for _, v := range vertex {
		vp := vectors.MatrixMultiply(v, clipMatrix)
		// Don't get confused here. I assume the divide leaves v.w alone.
		vp = vp.Div(v.W)
		dst = append(dst, vp)
	}

	/* TODO: Clipping here */
	//    Here, after the perspective divide, you perform Sutherland-Hodgeman clipping
	//    by checking if the x, y and z components are inside the range of [-w, w].
	//    One checks each vector component separately against each plane. Per-vertex
	//    data like colours, normals and texture coordinates need to be linearly
	//    interpolated for clipped edges to reflect the change. If the edge (v0,v1)
	//    is tested against the positive x plane, and v1 is outside, the interpolant
	//    becomes: (v1.x - w) / (v1.x - v0.x)
	//    I skip this stage all together to be brief.

	// and what does this do?
	// project to screen coordinates
	var sc []vectors.Vector
	for _, v := range dst {
		sc = append(sc, vectors.Vector{
			X: (v.X*float64(width))/(2.0*v.W) + halfWidth,
			Y: (v.Y*float64(height))/(2.0*v.W) + halfHeight,
		})
	}
	return sc
}

// SetupClipMatrix is the interesting stuff
func SetupClipMatrix(fov, aspectRatio float64) vectors.Matrix {
	f := 1.0 / math.Tan(fov*0.5)
	m := vectors.Matrix{}
	m.Data[0] = f * aspectRatio
	m.Data[5] = f
	m.Data[10] = (FAR_Z + NEAR_Z) / (FAR_Z - NEAR_Z)
	m.Data[11] = 1.0 // this 'plugs' the old z into w
	m.Data[14] = (2.0 * NEAR_Z * FAR_Z) / (NEAR_Z - FAR_Z)
	m.Data[15] = 0.0
	return m
}

func AsPNG4(name string, fill bool, vertices []vectors.Vector) error {
	dc := NewContext(width, height)

	r := 0.333 // rand.Float64()
	g := 0.333 // rand.Float64()
	b := 0.333 // rand.Float64()
	a := 1.0   // rand.Float64()*0.5 + 0.5

	dc.SetRGBA(r, g, b, a)
	for i, v := range vertices {
		if i == 0 {
			dc.MoveTo(v.X, v.Y)
		} else {
			dc.LineTo(v.X, v.Y)
		}
	}
	if fill {
		dc.StrokePreserve()
		dc.SetRGBA(1, 1, 1, 1)
		dc.Fill()
	} else {
		dc.Stroke()
	}

	return dc.SavePNG(name)
}
