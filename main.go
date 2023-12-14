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
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, k := range g.pressedKeys {
		switch k {
		case ebiten.KeyArrowUp:
			g.ship.OnUp()
		case ebiten.KeyArrowDown:
			g.ship.OnDown()
		case ebiten.KeyArrowLeft:
			g.ship.OnLeft()
		case ebiten.KeyArrowRight:
			g.ship.OnRight()
		}
	}
	g.ship.Drift()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ship: [%v,%v]\n\n", g.ship.Pos.X, g.ship.Pos.Y), int(g.ship.Pos.X), int(g.ship.Pos.Y))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ship: %v[%v,%v]", g.ship.Heading, g.ship.Speed.X, g.ship.Speed.Y), int(g.ship.Pos.X), int(g.ship.Pos.Y))

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
		WithSpeed(1, 0),
		WithMaxVelocity(2.5),
	}

	if err := ebiten.RunGame(NewGame(shipOptions...)); err != nil {
		log.Fatal(err)
	}
}
