package main

import "math"

type Vector struct {
	x, y, z, w float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{x: x, y: y, z: z, w: 1}
}

// Length assumes proper operator overloads here, with vectors and scalars
func (v Vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) Unit() Vector {
	const epsilon = 1e-6
	mag := v.Length()
	if mag < epsilon {
		panic("out of range")
	}
	return v.div(mag)
}

func (v Vector) div(div float64) Vector {
	// avoid divisions: they are much more costly than multiplications
	return v.mul(1 / div)
}

func (v Vector) dot(v2 Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v Vector) mul(fac float64) Vector {
	return Vector{x: v.x * fac, y: v.y * fac, z: v.z * fac, w: v.w}
}
