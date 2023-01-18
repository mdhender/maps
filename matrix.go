package main

type Matrix struct {
	data [16]float64
}

func (m Matrix) of(i int) float64 {
	if i < 0 || i > 15 {
		panic("index out of range")
	}
	return m.data[i]
}

func (m Matrix) mul(m2 Matrix) (dst Matrix) {
	for y := 0; y < 4; y++ {
		col := y * 4
		for x := 0; x < 4; x++ {
			for i := 0; i < 4; i++ {
				dst.data[x+col] += m2.data[i+col] * m.data[x+i*4]
			}
		}
	}
	return dst
}

func Identity() Matrix {
	m := Matrix{}
	m.data[0] = 1
	m.data[5] = 1
	m.data[10] = 1
	m.data[15] = 1
	return m
}
