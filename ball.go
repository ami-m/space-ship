package main

const (
	BallMaxVelocity = 50
	BallMinVelocity = 0
)

type Ball struct {
	PosX   float64
	PosY   float64
	SpeedX float64
	SpeedY float64
	Radius float64
}

type BallOption func(ball *Ball)

func WithPos(x, y float64) BallOption {
	return func(ball *Ball) {
		ball.PosX = x
		ball.PosY = y
	}
}

func WithSpeed(x, y float64) BallOption {
	return func(ball *Ball) {
		ball.SpeedX = x
		ball.SpeedY = y
	}
}

func NewBall(opts ...BallOption) *Ball {
	defaultRadius := 16.0
	res := Ball{
		Radius: defaultRadius,
		PosX:   defaultRadius,
		PosY:   defaultRadius,
	}
	for _, opt := range opts {
		opt(&res)
	}
	return &res
}

func (b *Ball) Drift() {
	b.PosX += b.SpeedX
	b.PosY += b.SpeedY
	b.OnWallCollision()
}

func (b *Ball) OnWallCollision() {
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

func (b *Ball) OnUp() {
	b.PosY -= 1
}

func (b *Ball) OnDown() {
	b.PosY += 1
}

func (b *Ball) OnLeft() {
	b.PosX -= 1
}

func (b *Ball) OnRight() {
	b.PosX += 1
}
