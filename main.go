package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenTitle    = "Pupuruneko Boids Simulation"
	screenWidth    = 640
	screenHeight   = 360
	screenScale    = 2
	boidCount      = 500
)

var (
	green    = color.RGBA{10, 255, 50, 255}
	boids    [boidCount]*Boid
	boidsMap [screenWidth + 1][screenHeight + 1]int
)

func main() {
	// Initialize boids map
	for i, row := range boidsMap {
		for j := range row {
			boidsMap[i][j] = -1
		}
	}

	// Initialize boids
	for i := 0; i < boidCount; i++ {
		CreateBoid(i)
	}

	// Render the simulation
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
