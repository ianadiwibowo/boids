package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector
	velocity Vector
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

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(renderDelay * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)

	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x <= 0 {
		b.velocity = Vector{
			x: -b.velocity.x,
			y: b.velocity.y,
		}
	}
	if next.y >= screenHeight || next.y <= 0 {
		b.velocity = Vector{
			x: b.velocity.x,
			y: -b.velocity.y,
		}
	}
}
