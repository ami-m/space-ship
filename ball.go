package main

const (
	BallMaxVelocity = 50
	BallMinVelocity = 0
)

type Ball struct {
	PosX   int
	PosY   int
	SpeedX int
	SpeedY int
	Radius int
}

type BallOption func(ball *Ball)

func WithPos(x, y int) BallOption {
	return func(ball *Ball) {
		ball.PosX = x
		ball.PosY = y
	}
}

func WithSpeed(x, y int) BallOption {
	return func(ball *Ball) {
		ball.SpeedX = x
		ball.SpeedY = y
	}
}

func NewBall(opts ...BallOption) *Ball {
	defaultRadius := 16
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

func (b *Ball) OnUp(step int) {
	b.PosY -= step
}

func (b *Ball) OnDown(step int) {
	b.PosY += step
}

func (b *Ball) OnLeft(step int) {
	b.PosX -= step
}

func (b *Ball) OnRight(step int) {
	b.PosX += step
}
