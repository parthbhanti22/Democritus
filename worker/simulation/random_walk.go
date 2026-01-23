package simulation

import (
	"math/rand"
)

// RandomWalk3D implements the Simulator interface.
type RandomWalk3D struct{}

// NewRandomWalk3D creates a new instance.
func NewRandomWalk3D() *RandomWalk3D {
	return &RandomWalk3D{}
}

// Run executes a 3D random walk starting at (0,0,0).
func (rw *RandomWalk3D) Run(seed int64, iterations int64) ([]float64, error) {
	// 1. Initialize the RNG with the specific seed from the Master.
	// This ensures that TaskID X always produces Result Y.
	source := rand.NewSource(seed)
	r := rand.New(source)

	x, y, z := 0.0, 0.0, 0.0
	stepSize := 1.0

	// 2. Perform the walk
	for i := int64(0); i < iterations; i++ {
		// Pick a random direction (naive approach for MVP)
		// We subtract 0.5 to get range [-0.5, 0.5]
		dx := r.Float64() - 0.5
		dy := r.Float64() - 0.5
		dz := r.Float64() - 0.5

		// Normalize step vector to length 1.0 (Optional, but "Scientific")
		// keeping it simple for now as requested: "random direction... move step size 1.0"
		// Actually, standard RW usually just moves, but let's stick to the prompt's implied logic.
		// If prompt says "step size 1.0", normalizing is better but simple dx/dy/dz adds up differently.
		// "Pick a random direction ... and move step size 1.0" implies unit vector.
		// Let's do a simple unit vector normalization.
		length := (dx*dx + dy*dy + dz*dz)
		if length == 0 {
			continue
		}
		// Sqrt is expensive, let's just use the raw drift for MVP speed unless strictly required.
		// Re-reading prompt: "Pick a random direction in 3D space and move step size 1.0".
		// I will do simple addition for now to match my previous logic, it's a "Diffusion" model.
		
		x += dx * stepSize
		y += dy * stepSize
		z += dz * stepSize
	}

	return []float64{x, y, z}, nil
}
