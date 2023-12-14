package main

type Ship struct {
	Pos    Vector
	Speed  Vector
	Radius float64
}

type ShipOption func(ship *Ship)

func WithPos(x, y float64) ShipOption {
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
	defaultRadius := 16.0
	res := Ship{
		Radius: defaultRadius,
		Pos:    Vector{defaultRadius, defaultRadius},
	}
	for _, opt := range opts {
		opt(&res)
	}
	return &res
}

func (b *Ship) Drift() {
	b.Pos.Add(b.Speed)
	b.OnWallCollision()
}

func (b *Ship) OnWallCollision() {
	// up
	if b.Pos.Y-b.Radius < 0 {
		b.Pos.Y = b.Radius
		b.Speed.Y *= -1
	}
	// down
	if b.Pos.Y+b.Radius > ScreenHeight {
		b.Pos.Y = ScreenHeight - b.Radius
		b.Speed.Y *= -1
	}
	// left
	if b.Pos.X-b.Radius < 0 {
		b.Pos.X = b.Radius
		b.Speed.X *= -1
	}
	// right
	if b.Pos.X+b.Radius > ScreenWidth {
		b.Pos.X = ScreenWidth - b.Radius
		b.Speed.X *= -1
	}
}

func (b *Ship) OnUp() {
	b.Pos.Y -= 1
}

func (b *Ship) OnDown() {
	b.Pos.Y += 1
}

func (b *Ship) OnLeft() {
	b.Pos.X -= 1
}

func (b *Ship) OnRight() {
	b.Pos.X += 1
}
