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
	ball        *Ball
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, k := range g.pressedKeys {
		switch k {
		case ebiten.KeyArrowUp:
			g.ball.OnUp()
		case ebiten.KeyArrowDown:
			g.ball.OnDown()
		case ebiten.KeyArrowLeft:
			g.ball.OnLeft()
		case ebiten.KeyArrowRight:
			g.ball.OnRight()
		}
	}
	g.ball.Drift()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ball: [%v,%v]\n\n", g.ball.PosX, g.ball.PosY), int(g.ball.PosX), int(g.ball.PosY))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame(ballOptions ...BallOption) *Game {
	return &Game{
		pressedKeys: nil,
		ball:        NewBall(ballOptions...),
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	ballOptions := []BallOption{
		WithSpeed(1, 1),
	}

	if err := ebiten.RunGame(NewGame(ballOptions...)); err != nil {
		log.Fatal(err)
	}
}
