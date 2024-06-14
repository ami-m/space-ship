package main

import (
	"game/events"
	"game/utils/timer"
	"game/utils/vector"
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/solarlune/resolv"
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

const (
	MapUp    = "up"
	MapDown  = "down"
	MapLeft  = "left"
	MapRight = "right"
	MapFire  = "fire"
)

type ShipKeyMapping map[ebiten.Key]string

type Ship struct {
	Id                  int
	Mass                float64
	eventPublisher      events.Subject
	ResolvObj           *resolv.Object
	Pos                 vector.Vector
	Speed               vector.Vector
	Radius              float64
	Heading             float64
	TurnRate            float64
	MaxVelocity         float64
	CollisionElasticity float64
	Acceleration        float64
	MuzzleSpeed         float64
	FireRate            time.Duration
	nextShotTimer       *timer.Timer
	SpritePath          string
	KeyMappings         ShipKeyMapping

	Shots []*ShipShot
}

func NewShip(id int, opts ...ShipOption) *Ship {
	res := Ship{
		Id:                  id,
		Mass:                1,
		eventPublisher:      events.NewEventPublisher(),
		Pos:                 vector.Vector{X: DefaultRadius, Y: DefaultRadius},
		Radius:              DefaultRadius,
		TurnRate:            DefaultTurnRate,
		MaxVelocity:         DefaultMaxVelocity,
		CollisionElasticity: DefaultCollisionElasticity,
		Acceleration:        DefaultAcceleration,
		MuzzleSpeed:         DefaultMuzzleSpeed,
		FireRate:            DefaultFireRate,
		SpritePath:          DefaultSpritePath,
		KeyMappings: map[ebiten.Key]string{
			ebiten.KeyArrowUp:    MapUp,
			ebiten.KeyArrowDown:  MapDown,
			ebiten.KeyArrowLeft:  MapLeft,
			ebiten.KeyArrowRight: MapRight,
			ebiten.KeySpace:      MapFire,
		},
	}
	for _, opt := range opts {
		opt(&res)
	}

	res.nextShotTimer = timer.NewTimer(res.FireRate)

	// build the resolv object for the ship
	res.ResolvObj = resolv.NewObject(res.Pos.X, res.Pos.Y, res.Radius*2, res.Radius*2, "ship")
	res.ResolvObj.Data = &res

	return &res
}

func (s *Ship) Update() {
	s.nextShotTimer.Update()
	s.Pos.Add(&s.Speed)
	s.updateResolver()

	//s.handleWallsCollision()
	//s.handleShipCollision()
	s.limitSpeed()
}

func (s *Ship) updateResolver() {
	s.ResolvObj.X = s.Pos.X - s.ResolvObj.W/2
	s.ResolvObj.Y = s.Pos.Y - s.ResolvObj.H/2
	s.ResolvObj.Update()
}

func (s *Ship) UpdateFromResolver() {
	s.Pos.X = s.ResolvObj.X + s.ResolvObj.W/2
	s.Pos.Y = s.ResolvObj.Y + s.ResolvObj.H/2
}

func (s *Ship) handleShipCollision() {
	if collision := s.ResolvObj.Check(s.Speed.X, s.Speed.Y, "ship"); collision != nil {
		for _, otherObject := range collision.Objects {
			otherShip := otherObject.Data.(*Ship)
			newSpeed := s.calcSpeedPostShipCollision(otherShip)
			s.Speed.X = newSpeed.X
			s.Speed.Y = newSpeed.Y
			//if isSideCollision(s.ResolvObj, otherObject) {
			//	s.Speed.X *= -1
			//}
			//if isTopDownCollision(s.ResolvObj, otherObject) {
			//	s.Speed.Y *= -1
			//}
			s.updateResolver()
		}

	}
}

func (s *Ship) calcSpeedPostShipCollision(otherShip *Ship) *vector.Vector {
	// Determine the collision vector (line of impact)
	collisionVector := otherShip.Pos.Clone().Subtract(&s.Pos).Normalize()

	// Decompose the velocity vectors into parallel and perpendicular components
	v1Parallel := collisionVector.Clone().Multiply(s.Speed.Dot(collisionVector))
	v1Perpendicular := s.Speed.Clone().Subtract(v1Parallel)

	v2Parallel := collisionVector.Clone().Multiply(otherShip.Speed.Dot(collisionVector))
	//v2Perpendicular := otherShip.Speed.Clone().Subtract(v2Parallel)

	// Calculate the new parallel components using the 1D elastic collision equations
	m1, m2 := s.Mass, otherShip.Mass
	newV1Parallel := v1Parallel.Clone().Multiply(m1 - m2).Add(v2Parallel.Clone().Multiply(2 * m2)).Multiply(1 / (m1 + m2))
	//newV2Parallel := v2Parallel.Multiply(m2 - m1).Add(v1Parallel.Multiply(2 * m1)).Multiply(1 / (m1 + m2))

	// Combine the parallel and perpendicular components
	newSpeed := newV1Parallel.Clone().Add(v1Perpendicular)
	return newSpeed
}

func isSideCollision(a, b *resolv.Object) bool {
	return isClose(a.X, b.Right()) || isClose(a.Right(), b.X)
}
func isTopDownCollision(a, b *resolv.Object) bool {
	return isClose(a.Y, b.Bottom()) || isClose(a.Bottom(), b.Y)
}

func isClose(a, b float64) bool {
	return math.Abs(a-b) < 5
}

func (s *Ship) handleWallsCollision() {
	if collision := s.ResolvObj.Check(s.Speed.X, s.Speed.Y, "wall"); collision != nil {
		for _, wall := range collision.Objects {
			s.handleWallCollision(wall)
			s.updateResolver()
		}

	}
}

func (s *Ship) handleWallCollision(wall *resolv.Object) {
	if wall.HasTags("ceiling") {
		s.Pos.Y = wall.Bottom() + s.ResolvObj.H/2 + 1
		s.Speed.Y *= s.CollisionElasticity
	}

	if wall.HasTags("floor") {
		s.Pos.Y = wall.Y - s.ResolvObj.H/2 - 1
		s.Speed.Y *= s.CollisionElasticity
	}

	if wall.HasTags("leftWall") {
		s.Pos.X = wall.Right() + s.ResolvObj.W/2 + 1
		s.Speed.X *= s.CollisionElasticity
	}

	if wall.HasTags("rightWall") {
		s.Pos.X = wall.X - s.ResolvObj.W/2 - 1
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

func (s *Ship) OnHitByShot() {
	log.Info("hit by shot")
	s.eventPublisher.FireEvent("shipHitByShot", events.EventPayload{
		"shipId": s.Id,
	})
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

func (s *Ship) OnFire() {
	if !s.nextShotTimer.IsReady() {
		return
	}
	s.nextShotTimer.Reset()

	// shot speed by heading
	shotSpeed := vector.Vector{
		X: s.MuzzleSpeed * math.Sin(s.Heading*math.Pi/180),
		Y: -1 * s.MuzzleSpeed * math.Cos(s.Heading*math.Pi/180),
	}
	shotSpeed.Add(&s.Speed)

	shotPosition := s.Pos
	shotPosition.Add(&vector.Vector{
		X: s.Radius / 2 * math.Sin(s.Heading*math.Pi/180),
		Y: -1 * s.Radius / 2 * math.Cos(s.Heading*math.Pi/180),
	})

	s.eventPublisher.FireEvent("shipFired", events.EventPayload{
		"shipId": s.Id,
		"ship":   s,
		"shot":   NewShipShot(s, shotPosition, shotSpeed, s.Heading),
	})
}

func (s *Ship) OnEvent(eventName string, payload events.EventPayload) {
	switch eventName {
	case "keyPress":
		s.handleKeyPress(payload["key"].(ebiten.Key))
	}
}

func (s *Ship) GetID() int {
	return s.Id
}

func (s *Ship) handleKeyPress(key ebiten.Key) {
	switch s.KeyMappings[key] {
	case MapUp:
		s.OnSpeedUp()
	case MapDown:
		s.OnSlowDown()
	case MapLeft:
		s.OnRotateLeft()
	case MapRight:
		s.OnRotateRight()
	case MapFire:
		s.OnFire()
	}
}

func (s *Ship) AddListener(observer events.Observer, eventName string) {
	s.eventPublisher.AddListener(observer, eventName)
}

func (s *Ship) RemoveListener(observer events.Observer, eventName string) {
	s.eventPublisher.RemoveListener(observer, eventName)
}

func (s *Ship) AdjustResolverSize(w, h float64) {
	s.ResolvObj.W = w
	s.ResolvObj.H = h
}
