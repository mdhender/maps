package draw

import (
	"bytes"
	"fmt"
	"os"
)

func AsSVG(name string) error {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("<svg style='stroke:grey; fill:none; stroke-width:0.7' width='%d' height='%d' xmlns='http://www.w3.org/2000/svg'>\n", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if nan(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			buf.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	buf.WriteString("</svg>")

	return os.WriteFile(name, buf.Bytes(), 0666)
}
