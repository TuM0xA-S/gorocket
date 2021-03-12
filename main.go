package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newRocketImg() *ebiten.Image {
	return ebiten.NewImageFromImage(RocketImgBase)
}

func (g *Game) generateBarrier() {
	b := Barrier{
		y:     -BarrierH,
		gateX: float64(rand.Int()%(BarrierEndX-GateW-BarrierStartX+1) + GateW),
	}
	g.barriers = append(g.barriers, b)
}

func (g *Game) removeRedundantBarrier() bool {
	if len(g.barriers) > 0 && g.barriers[0].y > ScreenH {
		g.barriers = g.barriers[1:]
		return true
	}
	return false
}

func (g *Game) moveBarriers() {
	for i := range g.barriers {
		g.barriers[i].y += g.barrierSpeedY
	}
}

//Game represent "Rocket" game logical state
type Game struct {
	barriers          []Barrier
	barrierTimer      int
	rocketX           float64
	mode              int
	smashTimer        int
	rocketImg         *ebiten.Image
	score             int
	record            int
	recordUpdated     bool
	barrierSpeedY     float64
	lvl               int
	barriersGoneCount int
	barrierDist       float64
}

func (g *Game) playerRect() Rect {
	return Rect{
		g.rocketX, RocketY,
		RocketW, RocketH,
	}
}

func (g *Game) moveRocket() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.rocketX -= RocketSpeedX
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.rocketX += RocketSpeedX
	}

	if g.rocketX > ScreenW {
		g.rocketX -= ScreenW
	}

	if g.rocketX+RocketW < 0 {
		g.rocketX += ScreenW
	}
}

func (g *Game) barrierProcessing() bool {

	if g.barrierTimer == 0 {
		g.generateBarrier()
		g.barrierTimer = g.getTicksPerBarrier()
	}
	res := g.removeRedundantBarrier()
	g.moveBarriers()

	g.barrierTimer--
	return res
}

func (g *Game) getTicksPerBarrier() int {
	return int((g.barrierDist + BarrierH) / g.barrierSpeedY)
}

//Update ...
func (g *Game) Update() error {
	switch g.mode {
	case ModePlay:
		barrierGone := g.barrierProcessing()
		if g.rocketSmashed() {
			g.mode = ModeSmash
			break
		}
		if barrierGone {
			g.score += int(math.Pow(ScoreMultipler, float64(g.lvl-1)))
			g.barriersGoneCount++
			if g.barriersGoneCount%PointsPerNextLevel == 0 {
				if g.lvl < MaxLevel {
					g.lvl++
				}
				if g.barrierSpeedY < BarrierSpeedYMax {
					g.barrierSpeedY += BarrierSpeedDelta
				}
				if g.barrierDist < BarrierDistMax {
					g.barrierDist += BarrierDistDelta
				}
			}
		}

		g.moveRocket()

	case ModeSmash:
		if g.smashTimer == 0 {
			g.mode = ModeGameOver
			break
		}
		g.barrierProcessing()
		//rocket pic demolition
		for i := 0; i < RocketSmashSpeed; i++ {
			g.rocketImg.Set(rand.Int()%RocketImgW,
				rand.Int()%RocketImgH, ColorTransparent)
		}
		g.smashTimer--

	case ModeStart:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset()
			g.mode = ModePlay
		}

	case ModeGameOver:
		if g.score > g.record {
			g.record = g.score
			g.recordUpdated = true
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModeStart
		}
	}
	return nil
}

func (g *Game) reset() {
	g.rocketX = ScreenW/2 - RocketW/2
	g.barriers = nil
	g.barrierTimer = 0
	g.smashTimer = RocketSmashTime
	g.rocketImg = newRocketImg()
	g.score = 0
	g.recordUpdated = false
	g.barrierSpeedY = BarrierSpeedYStarter
	g.lvl = 1
	g.barriersGoneCount = 0
	g.barrierDist = BarrierDistInitital
}

func (g *Game) rocketSmashed() bool {
	if len(g.barriers) == 0 {
		return false
	}
	r1, r2 := g.barriers[0].toRects()
	return g.playerRect().collides(r1) || g.playerRect().collides(r2)
}

//Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenW, ScreenH
}

func main() {
	ebiten.SetWindowSize(WindowW, WindowH)
	ebiten.SetWindowTitle("Rocket")
	g := &Game{record: loadRecord()}
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
	saveRecord(g.record)
}
