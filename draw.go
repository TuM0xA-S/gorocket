package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// x, y is coordinates of upper left
func drawTextWithBackground(dst *ebiten.Image, txt string,
	fg, bg color.Color, font font.Face, x, y float64) {

	r := text.BoundString(font, txt)
	tw := float64(r.Max.X)
	th := float64(font.Metrics().Height.Round())
	dx, dy := float64(r.Min.X), float64(r.Min.Y)
	const magicOffset = 10
	drawRect(dst, Rect{x - magicOffset, y - magicOffset, tw + 2*magicOffset, th}, bg)
	X1 := x - dx
	Y1 := y - dy
	text.Draw(dst, txt, font, int(X1), int(Y1), fg)
}

func drawRect(screen *ebiten.Image, r Rect, clr color.Color) {
	ebitenutil.DrawRect(screen, r.x, r.y, r.w, r.h, clr)
}

func (g *Game) drawBarriers(screen *ebiten.Image) {
	screen.Fill(ColorSpace)
	for _, b := range g.barriers {
		r1, r2 := b.toRects()
		drawRect(screen, r1, ColorGray)
		drawRect(screen, r2, ColorGray)
	}
}

func (g *Game) drawRocket(screen *ebiten.Image) {
	opGenerator := func(delta float64) *ebiten.DrawImageOptions {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, float64(RocketH)/float64(RocketImgH))
		op.GeoM.Translate(g.rocketX+delta, RocketY)
		return op
	}

	screen.DrawImage(g.rocketImg, opGenerator(0))
	screen.DrawImage(g.rocketImg, opGenerator(ScreenW))
	screen.DrawImage(g.rocketImg, opGenerator(-ScreenW))

}

func getCentredPosForText(txt string) float64 {
	W := text.BoundString(FontBig, txt).Max.X
	X := ScreenW/2 - W/2
	return float64(X)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	// drawRect(screen, Rect{ScoreRectX, ScoreRectY, ScoreRectW, ScoreRectH}, ColorBlack)
	score := fmt.Sprint(g.score)
	scoreX := getCentredPosForText(score)
	// text.Draw(screen, score, FontBig, scoreX, ScoreY, ColorGold)
	drawTextWithBackground(screen, score, ColorGold, ColorBlack,
		FontBig, float64(scoreX), 10)
}

//Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case ModePlay, ModeSmash:
		g.drawBarriers(screen)
		g.drawRocket(screen)
		g.drawScore(screen)
		ebitenutil.DebugPrint(screen, fmt.Sprint("level: ", g.lvl))

	case ModeStart:
		screen.DrawImage(StartSplashImg, nil)
		text.Draw(screen, fmt.Sprint(g.record), FontSmaller, RecordX, RecordY, ColorGold)

	case ModeGameOver:
		g.drawBarriers(screen)
		g.drawRocket(screen)

		drawTextWithBackground(screen, GameOverText, ColorGold, ColorBlack,
			FontBig, float64(GameOverTextX), GameOverTextY)
		text := fmt.Sprintf("Score: %d", g.score)

		GameOverScoreX := getCentredPosForText(text)
		drawTextWithBackground(screen, text, ColorGold, ColorBlack, FontBig, GameOverScoreX, GameOverScoreY)

		if g.recordUpdated {
			drawTextWithBackground(screen, NewRecordText, ColorGreen, ColorBlack, FontBig, NewRecordTextX, NewRecordTextY)
		}
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprint(len(g.barriers)))
}
