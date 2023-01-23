package main

import (
	"math"
)

const M_PI = math.Pi
const NEAR_Z = 0.5
const FAR_Z = 2.5

func Vectors() []Vector {
	var v []Vector
	// front face
	v = append(v, NewVector(0, 0, 30))
	v = append(v, NewVector(16, 0, 30))
	v = append(v, NewVector(16, 10, 30))
	v = append(v, NewVector(8, 16, 30))
	v = append(v, NewVector(0, 10, 30))
	v = append(v, NewVector(0, 0, 30))

	// rear face
	v = append(v, NewVector(0, 0, 54))
	v = append(v, NewVector(16, 0, 54))
	v = append(v, NewVector(16, 10, 54))
	v = append(v, NewVector(8, 16, 54))
	v = append(v, NewVector(0, 10, 54))
	v = append(v, NewVector(0, 0, 54))

	//// beams
	//v = append(v, NewVector(0, 0, 30))
	//v = append(v, NewVector(0, 0, 54))
	//v = append(v, NewVector(16, 0, 30))
	//v = append(v, NewVector(16, 0, 54))

	return ProjectAndClip(100, 100, v)
}

func MatrixMultiply(v Vector, m Matrix) Vector {
	return Vector{
		x: v.x*m.data[0] + v.y*m.data[4] + v.z*m.data[8] + v.w*m.data[12],
		y: v.x*m.data[1] + v.y*m.data[5] + v.z*m.data[9] + v.w*m.data[13],
		z: v.x*m.data[2] + v.y*m.data[6] + v.z*m.data[10] + v.w*m.data[14],
		w: v.x*m.data[3] + v.y*m.data[7] + v.z*m.data[11] + v.w*m.data[15],
	}
}

func ProjectAndClip(width, height int, vertex []Vector) (dst []Vector) {
	halfWidth := float64(width) * 0.5
	halfHeight := float64(height) * 0.5
	aspect := float64(width) / float64(height)
	clipMatrix := SetupClipMatrix(60.0*(math.Pi/180.0), aspect)

	for _, v := range vertex {
		vp := MatrixMultiply(v, clipMatrix)
		// Don't get confused here. I assume the divide leaves v.w alone.
		vp = vp.div(v.w)
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
	var sc []Vector
	for _, v := range dst {
		sc = append(sc, Vector{
			x: (v.x*float64(width))/(2.0*v.w) + halfWidth,
			y: (v.y*float64(height))/(2.0*v.w) + halfHeight,
		})
	}
	return sc
}

// SetupClipMatrix is the interesting stuff
func SetupClipMatrix(fov, aspectRatio float64) Matrix {
	f := 1.0 / math.Tan(fov*0.5)
	m := Matrix{}
	m.data[0] = f * aspectRatio
	m.data[5] = f
	m.data[10] = (FAR_Z + NEAR_Z) / (FAR_Z - NEAR_Z)
	m.data[11] = 1.0 // this 'plugs' the old z into w
	m.data[14] = (2.0 * NEAR_Z * FAR_Z) / (NEAR_Z - FAR_Z)
	m.data[15] = 0.0
	return m
}

func AsPNG4(name string, fill bool, vertices []Vector) error {
	dc := NewContext(width, height)

	r := 0.333 // rand.Float64()
	g := 0.333 // rand.Float64()
	b := 0.333 // rand.Float64()
	a := 1.0   // rand.Float64()*0.5 + 0.5

	dc.SetRGBA(r, g, b, a)
	for i, v := range vertices {
		if i == 0 {
			dc.MoveTo(v.x, v.y)
		} else {
			dc.LineTo(v.x, v.y)
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
