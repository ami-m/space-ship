package main

import "time"

type ShipOption func(ship *Ship)

func WithAlternateKeyMapping(keyMapping ShipKeyMapping) ShipOption {
	return func(s *Ship) {
		s.KeyMappings = keyMapping
	}
}

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
