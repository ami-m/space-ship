package health

import (
	"game/vector"
	"image/color"
)

// WithTotalPoints sets the TotalPoints field of the Bar.
func WithTotalPoints(totalPoints int) BarOption {
	return func(b *Bar) {
		b.TotalPoints = totalPoints
	}
}

// WithRemainingPoints sets the RemainingPoints field of the Bar.
func WithRemainingPoints(remainingPoints int) BarOption {
	return func(b *Bar) {
		b.RemainingPoints = remainingPoints
	}
}

// WithBorderColor sets the BorderColor field of the Bar.
func WithBorderColor(BorderColor color.Color) BarOption {
	return func(b *Bar) {
		b.BorderColor = BorderColor
	}
}

// WithBorderWidth sets the BorderWidth field of the Bar.
func WithBorderWidth(borderWidth int) BarOption {
	return func(b *Bar) {
		b.BorderWidth = borderWidth
	}
}

// WithPosition sets the Position field of the Bar.
func WithPosition(position vector.Vector) BarOption {
	return func(b *Bar) {
		b.Position = position
	}
}

// WithWidth sets the Width field of the Bar.
func WithWidth(width float64) BarOption {
	return func(b *Bar) {
		b.Width = width
	}
}

// WithHeight sets the Height field of the Bar.
func WithHeight(height float64) BarOption {
	return func(b *Bar) {
		b.Height = height
	}
}
