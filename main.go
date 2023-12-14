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
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ship: [%v,%v]\n\n", g.ship.PosX, g.ship.PosY), int(g.ship.PosX), int(g.ship.PosY))

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
		WithSpeed(1, 1),
	}

	if err := ebiten.RunGame(NewGame(shipOptions...)); err != nil {
		log.Fatal(err)
	}
}
