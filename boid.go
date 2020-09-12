package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector
	velocity Vector
}

const (
	nextMoveDelay  = 5 // Milliseconds
	viewRadius     = 13
	adjustmentRate = 0.015
)

func CreateBoid(id int) {
	boid := &Boid{
		id:       id,
		position: randomPosition(),
		velocity: randomVelocity(),
	}
	boid.RegisterToBoidsList()

	go boid.Fly()
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

func (b *Boid) RegisterToBoidsList() {
	boids[b.id] = b
	b.UpdatePositionInBoidsMap()
}

func (b *Boid) UpdatePositionInBoidsMap() {
	boidsMap[int(b.position.x)][int(b.position.y)] = b.id
}

func (b *Boid) Fly() {
	for {
		b.MoveOneStep()
		time.Sleep(nextMoveDelay * time.Millisecond)
	}
}

func (b *Boid) MoveOneStep() {
	b.velocity = b.velocity.Add(b.CalculateAcceleration()).Limit(-1, 1)

	boidsMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	b.UpdatePositionInBoidsMap()

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
	averageVelocity := Vector{0, 0}
	otherBoidsWithinViewRadiusCount := 0

	for i := math.Max(b.position.x-viewRadius, 0); i <= math.Min(b.position.x+viewRadius, screenWidth); i++ {
		for j := math.Max(b.position.y-viewRadius, 0); j <= math.Min(b.position.y+viewRadius, screenHeight); j++ {
			otherBoidID := boidsMap[int(i)][int(j)]
			if otherBoidID == -1 || otherBoidID == b.id {
				continue
			}
			otherBoidsWithinViewRadiusCount++
			averageVelocity = averageVelocity.Add(boids[otherBoidID].velocity)
		}
	}

	acceleration := Vector{0, 0}
	if (otherBoidsWithinViewRadiusCount) > 0 {
		averageVelocity = averageVelocity.DivideByValue(float64(otherBoidsWithinViewRadiusCount))
		acceleration = averageVelocity.Subtract(b.velocity).MultiplyByValue(adjustmentRate)
	}

	return acceleration
}
