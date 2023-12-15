package main

import (
	"math"
	"time"
)

const (
	DefaultMaxVelocity         float64 = 1
	DefaultCollisionElasticity float64 = -0.8
	DefaultTurnRate            float64 = 3
	DefaultRadius              float64 = 16
	DefaultAcceleration        float64 = 0.2
	DefaultMuzzleSpeed         float64 = 1.5
	DefaultFireRate                    = 200 * time.Millisecond
	DefaultSpritePath          string  = "assets/theme1/PNG/playerShip1_blue.png"
)

type Ship struct {
	Pos                 Vector
	Speed               Vector
	Radius              float64
	Heading             float64
	TurnRate            float64
	MaxVelocity         float64
	CollisionElasticity float64
	Acceleration        float64
	MuzzleSpeed         float64
	FireRate            time.Duration
	LastShotFiredAt     time.Time
	SpritePath          string

	Shots []*ShipShot
}

type ShipOption func(ship *Ship)

func WithHeading(heading float64) ShipOption {
	return func(ship *Ship) {
		ship.Heading = heading
	}
}

func WithSpritePath(path string) ShipOption {
	return func(ship *Ship) {
		ship.SpritePath = path
	}
}

func WithTurnRate(r float64) ShipOption {
	return func(ship *Ship) {
		ship.TurnRate = r
	}
}

func WithFireRate(r time.Duration) ShipOption {
	return func(ship *Ship) {
		ship.FireRate = r
	}
}

func WithMaxVelocity(m float64) ShipOption {
	return func(ship *Ship) {
		ship.MaxVelocity = m
	}
}

func WithPosition(x, y float64) ShipOption {
	return func(ship *Ship) {
		ship.Pos.X = x
		ship.Pos.Y = y
	}
}

func WithSpeed(x, y float64) ShipOption {
	return func(ship *Ship) {
		ship.Speed.X = x
		ship.Speed.Y = y
	}
}

func NewShip(opts ...ShipOption) *Ship {
	res := Ship{
		Pos:                 Vector{DefaultRadius, DefaultRadius},
		Radius:              DefaultRadius,
		TurnRate:            DefaultTurnRate,
		MaxVelocity:         DefaultMaxVelocity,
		CollisionElasticity: DefaultCollisionElasticity,
		Acceleration:        DefaultAcceleration,
		MuzzleSpeed:         DefaultMuzzleSpeed,
		FireRate:            DefaultFireRate,
		SpritePath:          DefaultSpritePath,
	}
	for _, opt := range opts {
		opt(&res)
	}
	return &res
}

func (s *Ship) Update() {
	s.Pos.Add(s.Speed)
	s.handleWallCollision()
	s.limitSpeed()
}

func (s *Ship) handleWallCollision() {
	// up
	if s.Pos.Y-s.Radius < 0 {
		s.Pos.Y = s.Radius
		s.Speed.Y *= s.CollisionElasticity
	}
	// down
	if s.Pos.Y+s.Radius > ScreenHeight {
		s.Pos.Y = ScreenHeight - s.Radius
		s.Speed.Y *= s.CollisionElasticity
	}
	// left
	if s.Pos.X-s.Radius < 0 {
		s.Pos.X = s.Radius
		s.Speed.X *= s.CollisionElasticity
	}
	// right
	if s.Pos.X+s.Radius > ScreenWidth {
		s.Pos.X = ScreenWidth - s.Radius
		s.Speed.X *= s.CollisionElasticity
	}
}

func (s *Ship) OnSpeedUp() {
	xPart, yPart := math.Sincos(s.Heading * math.Pi / 180.0)
	s.Speed.X += xPart * s.Acceleration
	s.Speed.Y -= yPart * s.Acceleration
}

func (s *Ship) limitSpeed() {
	if s.Speed.X > s.MaxVelocity {
		s.Speed.X = s.MaxVelocity
	}
	if s.Speed.Y > s.MaxVelocity {
		s.Speed.Y = s.MaxVelocity
	}
}

func (s *Ship) OnSlowDown() {
	s.Speed.X = s.Speed.X / 2
	s.Speed.Y = s.Speed.Y / 2
}

func (s *Ship) OnRotateLeft() {
	s.Heading -= s.TurnRate
	s.Heading = mod360(s.Heading)
}

func (s *Ship) OnRotateRight() {
	s.Heading += s.TurnRate
	s.Heading = mod360(s.Heading)
}

func mod360(n float64) float64 {
	for n < 0 {
		n += 360
	}
	for n >= 360 {
		n -= 360
	}
	return n
}

func (s *Ship) OnFire() *ShipShot {
	if time.Now().Sub(s.LastShotFiredAt) < s.FireRate {
		return nil
	}

	s.LastShotFiredAt = time.Now()

	// shot speed by heading
	shotSpeed := Vector{
		X: s.MuzzleSpeed * math.Sin(s.Heading*math.Pi/180),
		Y: -1 * s.MuzzleSpeed * math.Cos(s.Heading*math.Pi/180),
	}
	shotSpeed.Add(s.Speed)
	return NewShipShot(s, s.Pos, shotSpeed, s.Heading)
}
