package vectors

import "math"

const NEAR_Z = 0.5
const FAR_Z = 2.5

// SetupClip is the interesting stuff
func SetupClip(fov, aspectRatio float64) Matrix {
	f := 1.0 / math.Tan(fov*0.5)
	m := Matrix{}
	m.Data[0] = f * aspectRatio
	m.Data[5] = f
	m.Data[10] = (FAR_Z + NEAR_Z) / (FAR_Z - NEAR_Z)
	m.Data[11] = 1.0 // this 'plugs' the old z into w
	m.Data[14] = (2.0 * NEAR_Z * FAR_Z) / (NEAR_Z - FAR_Z)
	m.Data[15] = 0.0
	return m
}
