package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

func DrawShot(screen *ebiten.Image, s *ShipShot) {
	clr := s.OwningShip.Color

	x0 := s.Position.X
	y0 := s.Position.Y
	x1 := x0 + s.Radius*math.Sin(s.Heading*math.Pi/180)
	y1 := y0 - s.Radius*math.Cos(s.Heading*math.Pi/180)
	vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 1, clr, false)
}
