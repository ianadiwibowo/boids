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

const (
	nextMoveDelay = 5 // Milliseconds
)

func CreateBoid(id int) {
	boid := &Boid{
		id:       id,
		position: randomPosition(),
		velocity: randomVelocity(),
	}
	boid.UpdateBoidsList()
	boid.UpdateBoidsMap()

	go boid.Start()
}

func randomPosition() Vector {
	return Vector{
		x: rand.Float64() * screenWidth,
		y: rand.Float64() * screenHeight,
	}
}

func randomVelocity() Vector {
	return Vector{
		x: rand.Float64()*2 - 1.0,
		y: rand.Float64()*2 - 1.0,
	}
}

func (b *Boid) UpdateBoidsList() {
	boids[b.id] = b
}

func (b *Boid) UpdateBoidsMap() {
	boidsMap[int(b.position.x)][int(b.position.y)] = b.id
}

func (b *Boid) Start() {
	for {
		b.MoveOne()
		time.Sleep(nextMoveDelay * time.Millisecond)
	}
}

func (b *Boid) MoveOne() {
	b.velocity = b.velocity.Add(b.CalculateAcceleration()).Limit(-1, 1)

	boidsMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	b.UpdateBoidsMap()

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

func (b *Boid) CalculateAcceleration() Vector {
	return Vector{
		x: 0,
		y: 0,
	}
}
