package vectors

import "math"

type Vector struct {
	X, Y, Z, W float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z, W: 1}
}

// Length assumes proper operator overloads here, with vectors and scalars
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Unit() Vector {
	const epsilon = 1e-6
	mag := v.Length()
	if mag < epsilon {
		panic("out of range")
	}
	return v.Div(mag)
}

func (v Vector) Div(div float64) Vector {
	// avoid divisions: they are much more costly than multiplications
	return v.mul(1 / div)
}

func (v Vector) dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vector) mul(fac float64) Vector {
	return Vector{X: v.X * fac, Y: v.Y * fac, Z: v.Z * fac, W: v.W}
}
