package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
	"log"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	WindowWidth  = 640
	WindowHeight = 480
)

type Game struct {
	space       *resolv.Space
	pressedKeys []ebiten.Key
	ship1       *Ship
	ship2       *Ship
	Shots       []*ShipShot
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, k := range g.pressedKeys {
		switch k {
		case ebiten.KeyArrowUp:
			g.ship1.OnSpeedUp()
		case ebiten.KeyArrowDown:
			g.ship1.OnSlowDown()
		case ebiten.KeyArrowLeft:
			g.ship1.OnRotateLeft()
		case ebiten.KeyArrowRight:
			g.ship1.OnRotateRight()
		case ebiten.KeySpace:
			if shot := g.ship1.OnFire(); shot != nil {
				g.Shots = append(g.Shots, shot)
			}
		}
	}

	for _, k := range g.pressedKeys {
		switch k {
		case ebiten.KeyW:
			g.ship2.OnSpeedUp()
		case ebiten.KeyS:
			g.ship2.OnSlowDown()
		case ebiten.KeyA:
			g.ship2.OnRotateLeft()
		case ebiten.KeyD:
			g.ship2.OnRotateRight()
		case ebiten.KeyShiftLeft:
			if shot := g.ship2.OnFire(); shot != nil {
				g.Shots = append(g.Shots, shot)
			}
		}
	}

	g.ship1.Update()
	g.ship2.Update()

	g.removeOffScreenShots()
	for _, shot := range g.Shots {
		shot.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawShip(screen, g.ship1)
	DrawShip(screen, g.ship2)

	for _, shot := range g.Shots {
		DrawShot(screen, shot)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame(ship1, ship2 *Ship) *Game {
	res := Game{
		space:       resolv.NewSpace(ScreenWidth, ScreenHeight, 16, 16),
		pressedKeys: nil,
		ship1:       ship1,
		ship2:       ship2,
	}

	// space boundaries
	res.space.Add(
		resolv.NewObject(0, 0, ScreenWidth, 16, "wall", "ceiling"),
		resolv.NewObject(0, ScreenHeight-16, ScreenWidth, 16, "wall", "floor"),
		resolv.NewObject(0, 16, 16, ScreenHeight-32, "wall", "leftWall"),
		resolv.NewObject(ScreenWidth-16, 16, 16, ScreenHeight-32, "wall", "rightWall"),
	)
	res.space.Add(ship1.ResolvObj)
	res.space.Add(ship2.ResolvObj)

	return &res
}

func main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	ship1 := NewShip(WithSpritePath("assets/theme1/PNG/playerShip1_green.png"))
	ship2 := NewShip(
		WithPosition(100, 100),
		WithSpritePath("assets/theme1/PNG/playerShip3_blue.png"))

	if err := ebiten.RunGame(NewGame(ship1, ship2)); err != nil {
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
