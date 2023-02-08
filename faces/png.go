package faces

import "github.com/fogleman/gg"

// ToPNG writes the faces as a set of lines.
func (f Faces) ToPNG(dc *gg.Context) {
	for _, face := range f {
		dc.MoveTo(face[0].X, face[0].Y)
		dc.LineTo(face[1].X, face[1].Y)
		dc.LineTo(face[2].X, face[2].Y)
		dc.Stroke()
	}
}
