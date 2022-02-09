package main

import (
	"math"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	MIN_BRANCH_SIZE     = 1
	PI                  = 3.14159265359
	WIN_HEIGHT          = 600
	WIN_WIDTH           = 600
	FIRST_BRANCH_HEIGHT = 200
	ANGLE               = 60
)

func main() {
	_, renderer, err := sdl.CreateWindowAndRenderer(WIN_WIDTH, WIN_HEIGHT, sdl.WINDOW_BORDERLESS)
	if err != nil {
		panic(err)
	}
	changeAngle := int32(0)
	done := false
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	for !done {
		renderer.SetDrawColor(0xF, 0xF, 0xF, 0x0F)
		renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: WIN_WIDTH, H: WIN_HEIGHT})
		branch(0, FIRST_BRANCH_HEIGHT, WIN_WIDTH/2, WIN_HEIGHT, 0, changeAngle, 2, renderer)
		renderer.Present()
		changeAngle++
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				done = true
			}
		}
	}
}

func branch(depth, size, x, y, angle, angleInc int32, numberOfBranches int, renderer *sdl.Renderer) {
	if size <= MIN_BRANCH_SIZE {
		return
	}
	offsetX, offsetY := rotatePos(0, float64(-size), float64(angle))
	finalX := offsetX + float64(x)
	finalY := offsetY + float64(y)
	r, g, b := uint8(0xF0), uint8(0xF6), uint8(0xF0)

	gfx.ThickLineRGBA(renderer, x, y, int32(finalX), int32(finalY), 4, r, g, b, 0x0f)
	for i := -numberOfBranches / 2; i <= numberOfBranches/2; i++ {
		if i != 0 {
			branch(depth+1, int32(float32(size)*0.67), int32(finalX), int32(finalY), angle+angleInc*int32(i), angleInc, numberOfBranches, renderer)
		}
	}
}

func rotatePos(x, y, angle float64) (newX, newY float64) {
	angleInRadians := float32(angle) * (PI / 180.0)
	newX = (math.Cos(float64(angleInRadians)) * float64(x)) - (math.Sin(float64(angleInRadians)) * float64(y))
	newY = (math.Cos(float64(angleInRadians)) * float64(y)) + (math.Sin(float64(angleInRadians)) * float64(x))
	return
}
