package main

import "time"

type Boid struct {
	id       int
	position Vector
	velocity Vector
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
}
