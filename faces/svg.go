package faces

import (
	"bytes"
	"fmt"
)

// ToSVG writes the faces as a set of polygons.
func (f Faces) ToSVG(buf *bytes.Buffer, height float64) {
	for _, face := range f {
		face.ToSVG(buf, height)
	}
}

// ToSVG writes the face as a polygon.
func (f Face) ToSVG(buf *bytes.Buffer, height float64) {
	buf.WriteString("<polygon points='")
	for _, v := range f {
		buf.WriteString(fmt.Sprintf("%g,%g ", v.X, height-v.Y))
	}
	buf.WriteString("'/>\n")
}
