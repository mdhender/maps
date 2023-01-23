package vectors

func MatrixMultiply(v Vector, m Matrix) Vector {
	return Vector{
		X: v.X*m.Data[0] + v.Y*m.Data[4] + v.Z*m.Data[8] + v.W*m.Data[12],
		Y: v.X*m.Data[1] + v.Y*m.Data[5] + v.Z*m.Data[9] + v.W*m.Data[13],
		Z: v.X*m.Data[2] + v.Y*m.Data[6] + v.Z*m.Data[10] + v.W*m.Data[14],
		W: v.X*m.Data[3] + v.Y*m.Data[7] + v.Z*m.Data[11] + v.W*m.Data[15],
	}
}
