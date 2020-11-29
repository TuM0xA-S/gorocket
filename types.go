package main

import "math"

//Barrier represent ===== ===== this
type Barrier struct {
	y     float64
	gateX float64
}

//Rect represent rectangle
type Rect struct {
	x, y float64
	w, h float64
}

func (b Barrier) toRects() (Rect, Rect) {
	const barrierW = ScreenW - GateW
	return Rect{b.gateX, b.y, barrierW, BarrierH},
		Rect{b.gateX - ScreenW, b.y, barrierW, BarrierH}
}

func (r Rect) collides(r1 Rect) bool {
	x := math.Max(r.x, r1.x)
	x1 := math.Min(r.x+r.w, r1.x+r1.w)
	y := math.Max(r.y, r1.y)
	y1 := math.Min(r.y+r.h, r1.y+r1.h)

	const gap = 3
	return gap < x1-x && gap < y1-y
}
