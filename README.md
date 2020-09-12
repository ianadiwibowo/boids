# Boids

A simulation of birds flocking beautifully in the sky, with goroutines and Ebiten.

## About Boids

Boids is an artificial life simulation which mimics the birds flocking behavior. It was created by Craig Reynolds who published this topic in ACM SIGGRAPPH 1987 conference. The term boid itself stands for "bird-android" or "bird-oid object".

Boids simulation shows an emergent behavior. Each boid interacts with other individual boids following a set of rules: separation, alignment, and cohesion.

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/cd/Starling_flock_with_nearby_predator.jpg/872px-Starling_flock_with_nearby_predator.jpg" height="300" alt="Birds Flocking 1"> <img src="https://upload.wikimedia.org/wikipedia/commons/9/92/Sort_sol_pdfnet.jpg" height="300" alt="Birds Flocking 1">

## Instruction

```bash
# Clone this repository
git clone https://github.com/ianadiwibowo/boids.git
cd boids

# Setup the dependencies
go mod download

# Run the simulation
go run github.com/ianadiwibowo/boids
```

## References

- https://en.wikipedia.org/wiki/Boids
- https://gamedevelopment.tutsplus.com/tutorials/3-simple-rules-of-flocking-behaviors-alignment-cohesion-and-separation--gamedev-3444
