package main

import "time"

type ShipShot struct {
	OwningShip *Ship
	FiredAt    time.Time
	Position   Vector
	Speed      Vector
	Heading    float64
	Radius     float64
	OffScreen  bool
}

func NewShipShot(ship *Ship, pos Vector, speed Vector, heading float64) *ShipShot {
	return &ShipShot{
		OwningShip: ship,
		FiredAt:    time.Now(),
		Position:   pos,
		Speed:      speed,
		Heading:    heading,
		Radius:     5,
		OffScreen:  false,
	}
}

func (s *ShipShot) Update() {
	s.Position.Add(s.Speed)
	s.handleWallCollision()
}

func (s *ShipShot) handleWallCollision() {
	// up
	if s.Position.Y-s.Radius < 0 {
		s.OffScreen = true
	}
	// down
	if s.Position.Y+s.Radius > ScreenHeight {
		s.OffScreen = true
	}
	// left
	if s.Position.X-s.Radius < 0 {
		s.OffScreen = true
	}
	// right
	if s.Position.X+s.Radius > ScreenWidth {
		s.OffScreen = true
	}
}
