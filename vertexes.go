package main

import (
	"github.com/mdhender/maps/faces"
)

// Vertexes returns the faces that compose the image
func Vertexes() faces.Faces {
	var f faces.Faces

	// front
	f = append(f, faces.New(0, 0, 30, 0, 10, 30, 16, 10, 30))
	f = append(f, faces.New(0, 0, 30, 16, 0, 30, 16, 10, 30))

	// rear
	f = append(f, faces.New(0, 0, 54, 0, 10, 54, 16, 10, 54))
	f = append(f, faces.New(0, 0, 54, 16, 0, 54, 16, 10, 54))

	// front top
	f = append(f, faces.New(0, 10, 30, 8, 16, 30, 16, 10, 30))

	// rear top
	f = append(f, faces.New(0, 10, 54, 8, 16, 54, 16, 10, 54))

	// left side
	f = append(f, faces.New(0, 0, 30, 0, 0, 54, 0, 10, 54))
	f = append(f, faces.New(0, 0, 30, 0, 10, 30, 0, 10, 54))

	// right side
	f = append(f, faces.New(16, 0, 30, 16, 0, 54, 16, 10, 54))
	f = append(f, faces.New(16, 0, 30, 16, 10, 30, 16, 10, 54))

	// left roof
	f = append(f, faces.New(0, 10, 30, 0, 10, 54, 8, 16, 54))
	f = append(f, faces.New(0, 10, 30, 8, 16, 30, 8, 16, 54))

	// right roof
	f = append(f, faces.New(16, 10, 30, 16, 10, 54, 8, 16, 54))
	f = append(f, faces.New(16, 10, 30, 8, 16, 30, 8, 16, 54))

	// bottom
	f = append(f, faces.New(0, 0, 30, 0, 0, 54, 16, 0, 54))
	f = append(f, faces.New(0, 0, 30, 16, 0, 30, 16, 0, 54))

	return f
}
