package main

import (
	"game/utils/vector"
	"github.com/solarlune/resolv"
	"time"
)

type ShipShot struct {
	ResolvObj  *resolv.Object
	OwningShip *Ship
	FiredAt    time.Time
	Pos        vector.Vector
	Speed      vector.Vector
	Heading    float64
	Radius     float64
	OffScreen  bool
}

func NewShipShot(ship *Ship, pos vector.Vector, speed vector.Vector, heading float64) *ShipShot {
	res := ShipShot{
		OwningShip: ship,
		FiredAt:    time.Now(),
		Pos:        pos,
		Speed:      speed,
		Heading:    heading,
		Radius:     5,
		OffScreen:  false,
	}

	// build the resolv object
	res.ResolvObj = resolv.NewObject(res.Pos.X, res.Pos.Y, res.Radius, res.Radius, "shipShot")
	res.ResolvObj.Data = &res

	return &res
}

func (s *ShipShot) Update() {
	s.Pos.Add(s.Speed)
	s.updateResolver()
	s.handleWallCollision()
	s.handleShipCollision()
}

func (s *ShipShot) updateResolver() {
	s.ResolvObj.X = s.Pos.X
	s.ResolvObj.Y = s.Pos.Y
	s.ResolvObj.Update()
}

func (s *ShipShot) handleWallCollision() {
	if collision := s.ResolvObj.Check(s.Speed.X, s.Speed.Y, "wall"); collision != nil {
		s.OffScreen = true
	}
}

func (s *ShipShot) handleShipCollision() {
	if collision := s.ResolvObj.Check(s.Speed.X, s.Speed.Y, "ship"); collision != nil {
		if ship, ok := collision.Objects[0].Data.(*Ship); ok {
			if ship == s.OwningShip {
				return
			}

			s.OffScreen = true
			ship.OnHitByShot()
		}

	}
}
