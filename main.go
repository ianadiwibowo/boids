package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 360
	screenScale  = 2
	screenTitle  = "Pupuruneko Boids Simulation"
	boidCount    = 500
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

func createBoid(id int) {
	boid := &Boid{
		id: id,
		position: Vector{
			x: rand.Float64() * screenWidth,
			y: rand.Float64() * screenHeight,
		},
		velocity: Vector{
			x: rand.Float64()*2 - 1.0,
			y: rand.Float64()*2 - 1.0,
		},
	}
	boids[id] = boid
	go boid.start()
}

func updateScreen(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for boid := range boids {
			// screen.Set(boid.position, y int, clr color.Color)
		}
	}
}
