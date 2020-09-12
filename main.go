package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenTitle  = "Pupuruneko Boids Simulation"
	screenWidth  = 640
	screenHeight = 360
	screenScale  = 2
	boidCount    = 500
	renderDelay  = 5
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	err := ebiten.Run(
		updateScreen,
		screenWidth,
		screenHeight,
		screenScale,
		screenTitle,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func updateScreen(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x), int(boid.position.y), green)
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		}
	}
	return nil
}
