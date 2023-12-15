package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

type Game struct {
	pressedKeys []ebiten.Key
	ship        *Ship
	Shots       []*ShipShot
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, k := range g.pressedKeys {
		switch k {
		case ebiten.KeyArrowUp:
			g.ship.OnSpeedUp()
		case ebiten.KeyArrowDown:
			g.ship.OnSlowDown()
		case ebiten.KeyArrowLeft:
			g.ship.OnRotateLeft()
		case ebiten.KeyArrowRight:
			g.ship.OnRotateRight()
		case ebiten.KeySpace:
			if shot := g.ship.OnFire(); shot != nil {
				g.Shots = append(g.Shots, shot)
			}
		}
	}

	g.ship.Update()

	g.removeOffScreenShots()
	for _, shot := range g.Shots {
		shot.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ship: %v[%v,%v]", g.ship.Heading, g.ship.Speed.X, g.ship.Speed.Y), int(g.ship.Pos.X), int(g.ship.Pos.Y))
	DrawShip(screen, g.ship)

	for _, shot := range g.Shots {
		DrawShot(screen, shot)
	}
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("shots: %d", len(g.Shots)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame(shipOptions ...ShipOption) *Game {
	return &Game{
		pressedKeys: nil,
		ship:        NewShip(shipOptions...),
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	shipOptions := []ShipOption{
		//WithSpeed(1, 0),
		WithMaxVelocity(1),
	}

	if err := ebiten.RunGame(NewGame(shipOptions...)); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) removeOffScreenShots() {
	shots := make([]*ShipShot, 0)
	for _, shot := range g.Shots {
		if !shot.OffScreen {
			shots = append(shots, shot)
		}
	}
	g.Shots = shots
}
