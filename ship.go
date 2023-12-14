package main

type Ship struct {
	PosX   float64
	PosY   float64
	SpeedX float64
	SpeedY float64
	Radius float64
}

type ShipOption func(ship *Ship)

func WithPos(x, y float64) ShipOption {
	return func(ship *Ship) {
		ship.PosX = x
		ship.PosY = y
	}
}

func WithSpeed(x, y float64) ShipOption {
	return func(ship *Ship) {
		ship.SpeedX = x
		ship.SpeedY = y
	}
}

func NewShip(opts ...ShipOption) *Ship {
	defaultRadius := 16.0
	res := Ship{
		Radius: defaultRadius,
		PosX:   defaultRadius,
		PosY:   defaultRadius,
	}
	for _, opt := range opts {
		opt(&res)
	}
	return &res
}

func (b *Ship) Drift() {
	b.PosX += b.SpeedX
	b.PosY += b.SpeedY
	b.OnWallCollision()
}

func (b *Ship) OnWallCollision() {
	// up
	if b.PosY-b.Radius < 0 {
		b.PosY = b.Radius
		b.SpeedY *= -1
	}
	// down
	if b.PosY+b.Radius > ScreenHeight {
		b.PosY = ScreenHeight - b.Radius
		b.SpeedY *= -1
	}
	// left
	if b.PosX-b.Radius < 0 {
		b.PosX = b.Radius
		b.SpeedX *= -1
	}
	// right
	if b.PosX+b.Radius > ScreenWidth {
		b.PosX = ScreenWidth - b.Radius
		b.SpeedX *= -1
	}
}

func (b *Ship) OnUp() {
	b.PosY -= 1
}

func (b *Ship) OnDown() {
	b.PosY += 1
}

func (b *Ship) OnLeft() {
	b.PosX -= 1
}

func (b *Ship) OnRight() {
	b.PosX += 1
}
