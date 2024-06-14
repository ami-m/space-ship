package vector

import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(v2 *Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector) Clone() *Vector {
	return &Vector{
		X: v.X,
		Y: v.Y,
	}
}

func (v *Vector) Subtract(other *Vector) *Vector {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v *Vector) Multiply(factor float64) *Vector {
	v.X *= factor
	v.Y *= factor
	return v
}

func (v *Vector) Dot(other *Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() *Vector {
	mag := v.Magnitude()
	if mag == 0 {
		v.X = 0
		v.Y = 0
	} else {
		v.X /= mag
		v.Y /= mag
	}

	return v
}
