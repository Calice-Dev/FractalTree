package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	MIN_BRANCH_SIZE     = 1
	PI                  = 3.14159265359
	WIN_HEIGHT          = 1200
	WIN_WIDTH           = 1200
	FIRST_BRANCH_HEIGHT = 400
	ANGLE               = 60
)

func main() {
	_, renderer, err := sdl.CreateWindowAndRenderer(WIN_WIDTH, WIN_HEIGHT, sdl.WINDOW_BORDERLESS)
	if err != nil {
		panic(err)
	}
	angle := 0
	for {
		renderer.SetDrawColor(0x00, 0x00, 0x0, 0xff)
		renderer.Clear()

		branch(0, FIRST_BRANCH_HEIGHT, WIN_WIDTH/2, WIN_HEIGHT, int32(angle), renderer)
		renderer.Present()
		//angle++

		//window.UpdateSurface()
	}
}

func branch(depth, size, x, y, angle int32, renderer *sdl.Renderer) {
	if size <= MIN_BRANCH_SIZE {
		return
	}
	//rand.Seed(0)
	offsetX, offsetY := rotatePos(0, float64(-size), float64(angle))
	finalX := offsetX + float64(x)
	finalY := offsetY + float64(y)
	renderer.SetDrawColor(uint8(depth), 0xff, 0xff, 0xff)
	renderer.DrawLine(x, y, int32(finalX), int32(finalY))
	branch(depth+1, int32(float32(size)*0.67), int32(finalX), int32(finalY), angle+ANGLE, renderer)

	branch(depth+1, int32(float32(size)*0.67), int32(finalX), int32(finalY), angle-ANGLE, renderer)

}

func rotatePos(x, y, angle float64) (newX, newY float64) {
	angleInRadians := float32(angle) * (PI / 180.0)
	newX = (math.Cos(float64(angleInRadians)) * float64(x)) - (math.Sin(float64(angleInRadians)) * float64(y))
	newY = (math.Cos(float64(angleInRadians)) * float64(y)) + (math.Sin(float64(angleInRadians)) * float64(x))
	return
}
