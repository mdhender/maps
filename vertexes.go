package main

import (
	"github.com/mdhender/maps/faces"
)

// MakeTriangle returns the faces that compose the image of a triangle.
func MakeTriangle() (f faces.Faces) {
	// front
	f = append(f, faces.New(
		0, 0, 0,
		30, 0, 0,
		0, 30, 0))

	return f
}

// MakeHouse returns the faces that compose the image of a house
func MakeHouse() (f faces.Faces) {
	// front
	f = append(f, faces.New(
		0, 0, 0,
		24, 0, 0,
		24, 24, 0))
	f = append(f, faces.New(
		0, 0, 0,
		24, 24, 0,
		0, 24, 0))

	// front top
	f = append(f, faces.New(
		0, 24, 0,
		24, 24, 0,
		12, 32, 0))

	// rear
	f = append(f, faces.New(
		0, 0, -24,
		24, 0, -24,
		24, 24, -24))
	f = append(f, faces.New(
		0, 0, -24,
		24, 24, -24,
		0, 24, -24))

	// rear top
	f = append(f, faces.New(
		0, 24, -24,
		24, 24, -24,
		12, 32, -24))

	// left side
	f = append(f, faces.New(
		0, 0, 0,
		0, 0, -24,
		0, 24, -24))
	f = append(f, faces.New(
		0, 0, 0,
		0, 24, -24,
		0, 24, 0))

	// right side
	f = append(f, faces.New(
		24, 0, 0,
		24, 0, -24,
		24, 24, -24))
	f = append(f, faces.New(
		24, 0, 0,
		24, 24, -24,
		24, 24, 0))

	// bottom
	f = append(f, faces.New(
		0, 0, 0,
		24, 0, 0,
		24, 0, -24))
	f = append(f, faces.New(
		0, 0, 0,
		24, 0, -24,
		0, 0, -24))

	return f
}
