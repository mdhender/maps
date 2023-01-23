package vectors

type Matrix struct {
	Data [16]float64
}

func (m Matrix) of(i int) float64 {
	if i < 0 || i > 15 {
		panic("index out of range")
	}
	return m.Data[i]
}

func (m Matrix) mul(m2 Matrix) (dst Matrix) {
	for y := 0; y < 4; y++ {
		col := y * 4
		for x := 0; x < 4; x++ {
			for i := 0; i < 4; i++ {
				dst.Data[x+col] += m2.Data[i+col] * m.Data[x+i*4]
			}
		}
	}
	return dst
}

func Identity() Matrix {
	m := Matrix{}
	m.Data[0] = 1
	m.Data[5] = 1
	m.Data[10] = 1
	m.Data[15] = 1
	return m
}
