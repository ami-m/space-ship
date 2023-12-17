package health

import (
	"game/utils/vector"
	"github.com/hajimehoshi/ebiten/v2"
	ebitenVec "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Bar struct {
	TotalPoints     int
	RemainingPoints int
	BorderColor     color.Color
	BorderWidth     int
	Position        vector.Vector
	Width           float64
	Height          float64
}

type BarOption func(b *Bar)

func NewBar(opts ...BarOption) *Bar {
	res := Bar{
		TotalPoints:     100,
		RemainingPoints: 100,
		Position: vector.Vector{
			X: 32,
			Y: 32,
		},
		Width:       160,
		Height:      16,
		BorderColor: color.RGBA{G: 0xff, A: 0xff},
		BorderWidth: 1,
	}

	for _, opt := range opts {
		opt(&res)
	}

	return &res
}

func (b *Bar) IncPoints(n int) {
	b.RemainingPoints += n
	if b.RemainingPoints < 0 {
		b.RemainingPoints = 0
	}
	if b.RemainingPoints > b.TotalPoints {
		b.RemainingPoints = b.TotalPoints
	}
}

func (b *Bar) Draw(screen *ebiten.Image) {
	// Calculate the current color based on the remaining points.
	currentColor := b.calcCurrentColor()

	// Calculate the width of the filled health bar based on the remaining points.
	barWidth := (b.Width - float64(2*b.BorderWidth)) * (float64(b.RemainingPoints) / float64(b.TotalPoints))

	// Draw the health bar frame.
	ebitenVec.StrokeRect(screen, float32(b.Position.X), float32(b.Position.Y), float32(b.Width), float32(b.Height), float32(b.BorderWidth), b.BorderColor, false)

	// Draw the filled health bar based on the current color.
	ebitenVec.DrawFilledRect(
		screen,
		float32(b.Position.X)+float32(b.BorderWidth),
		float32(b.Position.Y)+float32(b.BorderWidth),
		float32(barWidth),
		float32(b.Height)-float32(2*b.BorderWidth),
		currentColor,
		false)
}

func (b *Bar) calcCurrentColor() color.Color {
	return color.RGBA{R: 0xff, A: 0xff}
	// Calculate the ratio of RemainingPoints to TotalPoints as a value between 0 and 1.
	//ratio := float64(b.RemainingPoints) / float64(b.TotalPoints)
	//
	//// Calculate the color components based on the ratio.
	//// 100% (1.0) is green, 0% (0.0) is red, and 50% (0.5) is yellow.
	//r := uint8(math.Max(0, 255*(1.0-ratio-0.5)))
	//g := uint8(math.Max(0, 255*(ratio-0.5)))
	//brightness := uint8(math.Min(255, 255*math.Abs(0.5-ratio)*2))
	//
	//// Create and return the color.
	//return color.RGBA{R: r, G: g, B: 0, A: brightness}
}
