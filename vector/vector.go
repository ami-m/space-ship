package vector

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
