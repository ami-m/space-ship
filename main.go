package main

import (
	"fmt"
	"game/events"
	"game/health"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitenVec "github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
	"image/color"
	"log"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	WindowWidth  = 640
	WindowHeight = 480
)

type Game struct {
	space          *resolv.Space
	eventPublisher events.Subject
	pressedKeys    []ebiten.Key
	ship1          *Ship
	ship2          *Ship
	Shots          []*ShipShot
	bar1           *health.Bar
	bar2           *health.Bar
	ceiling        *resolv.Object
	floor          *resolv.Object
	leftWall       *resolv.Object
	rightWall      *resolv.Object
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	g.publishKeysPressed()

	g.ship1.Update()
	g.ship2.Update()

	g.removeOffScreenShots()
	for _, shot := range g.Shots {
		shot.Update()
	}

	return nil
}

func printResolvObjAt(screen *ebiten.Image, x, y int, obj *resolv.Object) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v at [%v,%v] W:%v H:%v", obj.Tags(), obj.X, obj.Y, obj.W, obj.H), x, y)
}

func drawResolvObject(screen *ebiten.Image, obj *resolv.Object) {
	ebitenVec.DrawFilledRect(screen, float32(obj.X), float32(obj.Y), float32(obj.W), float32(obj.H), color.RGBA{
		R: 0xff,
	}, false)
}

func (g *Game) Draw(screen *ebiten.Image) {
	//drawResolvObject(screen, g.ceiling)
	//drawResolvObject(screen, g.floor)
	//drawResolvObject(screen, g.leftWall)
	//drawResolvObject(screen, g.rightWall)
	//drawResolvObject(screen, g.ship1.ResolvObj)

	// draw ships
	DrawShip(screen, g.ship1)
	DrawShip(screen, g.ship2)

	// draw shots
	for _, shot := range g.Shots {
		DrawShot(screen, shot)
	}

	// draw health bars
	g.bar1.Draw(screen)
	g.bar2.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame(ship1, ship2 *Ship) *Game {
	bar1Options := []health.BarOption{
		health.WithPosition(5, 5),
		health.WithHeight(5),
		health.WithWidth(80),
	}
	bar2Options := []health.BarOption{
		health.WithPosition(ScreenWidth-80-2, 5),
		health.WithHeight(5),
		health.WithWidth(80),
	}

	res := Game{
		eventPublisher: events.NewEventPublisher(),
		space:          resolv.NewSpace(ScreenWidth, ScreenHeight, 16, 16),
		pressedKeys:    nil,
		ship1:          ship1,
		ship2:          ship2,
		bar1:           health.NewBar(bar1Options...),
		bar2:           health.NewBar(bar2Options...),
	}

	//TODO: make both game and ships depend on one central observer
	res.eventPublisher.AddListener(ship1, "keyPress")
	res.eventPublisher.AddListener(ship2, "keyPress")

	ship1.AddListener(&res, "shipFired")
	ship1.AddListener(&res, "shipHitByShot")
	ship2.AddListener(&res, "shipFired")
	ship2.AddListener(&res, "shipHitByShot")

	// space boundaries
	res.ceiling = resolv.NewObject(0, 16, ScreenWidth, 16, "wall", "ceiling")
	res.floor = resolv.NewObject(0, ScreenHeight-16, ScreenWidth, 16, "wall", "floor")
	res.leftWall = resolv.NewObject(0, 16, 16, ScreenHeight-32, "wall", "leftWall")
	res.rightWall = resolv.NewObject(ScreenWidth-16, 16, 16, ScreenHeight-32, "wall", "rightWall")
	res.space.Add(
		res.ceiling,
		res.floor,
		res.leftWall,
		res.rightWall,
	)
	res.space.Add(ship1.ResolvObj)
	res.space.Add(ship2.ResolvObj)

	return &res
}

func main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	// TODO: all the ship config should be moved elsewhere (from a file?)
	ship1 := NewShip(
		1,
		WithPosition(250, 150),
		WithSpritePath("assets/theme1/PNG/playerShip1_green.png"))
	ship2 := NewShip(
		2,
		WithPosition(50, 50),
		WithSpritePath("assets/theme1/PNG/playerShip3_blue.png"),
		WithAlternateKeyMapping(ShipKeyMapping{
			ebiten.KeyW:         MapUp,
			ebiten.KeyS:         MapDown,
			ebiten.KeyA:         MapLeft,
			ebiten.KeyD:         MapRight,
			ebiten.KeyShiftLeft: MapFire,
		}))

	if err := ebiten.RunGame(NewGame(ship1, ship2)); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) addShot(shot *ShipShot) {
	g.Shots = append(g.Shots, shot)
	g.space.Add(shot.ResolvObj)
}

func (g *Game) removeOffScreenShots() {
	shots := make([]*ShipShot, 0)
	for _, shot := range g.Shots {
		if !shot.OffScreen {
			shots = append(shots, shot)
		} else {
			g.space.Remove(shot.ResolvObj)
		}
	}
	g.Shots = shots
}

func (g *Game) publishKeysPressed() {
	for _, k := range g.pressedKeys {
		g.eventPublisher.FireEvent("keyPress", map[string]any{"key": k})
	}
}

func (g *Game) OnEvent(eventName string, payload events.EventPayload) {
	switch eventName {
	case "shipFired":
		shot := payload["shot"].(*ShipShot)
		g.addShot(shot)
	case "shipHitByShot":
		shipId := payload["shipId"].(int)
		if shipId == 1 {
			g.bar1.IncPoints(-1)
		} else {
			g.bar2.IncPoints(-1)
		}
	}
}

func (g *Game) GetID() int {
	return 0
}
