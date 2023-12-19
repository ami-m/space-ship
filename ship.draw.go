package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math"
)

//go:embed assets/*
var assets embed.FS

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func DrawShip(screen *ebiten.Image, s *Ship) {
	const scaleCorrection = 0.25
	var PlayerSprite = mustLoadImage(s.SpritePath)
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(scaleCorrection, scaleCorrection)

	width := PlayerSprite.Bounds().Dx()
	height := PlayerSprite.Bounds().Dy()
	halfW := float64(width) * scaleCorrection / 2
	halfH := float64(height) * scaleCorrection / 2

	s.AdjustResolverSize(halfW*2, halfH*2)

	op.GeoM.Translate(-halfW, -halfH)

	op.GeoM.Rotate(s.Heading * math.Pi / 180.0)
	op.GeoM.Translate(s.Pos.X, s.Pos.Y)

	screen.DrawImage(PlayerSprite, op)
}
