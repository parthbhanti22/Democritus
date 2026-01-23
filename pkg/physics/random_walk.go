package physics

import (
	"math/rand"
)

// RandomWalk3D implements the Simulation interface for a particle diffusion model.
type RandomWalk3D struct{}

// NewRandomWalk3D creates a new instance of the model.
func NewRandomWalk3D() *RandomWalk3D {
	return &RandomWalk3D{}
}

// Run executes a 3D random walk.
// It returns the final [x, y, z] coordinates.
func (rw *RandomWalk3D) Run(seed int64, iterations int64) ([]float64, error) {
	// 1. Create a Source using the provided seed.
	// This ensures that if we pass seed=12345, we ALWAYS get the same path.
	source := rand.NewSource(seed)
	r := rand.New(source)

	x, y, z := 0.0, 0.0, 0.0

	for i := int64(0); i < iterations; i++ {
		// Randomly choose a direction in 3D space.
		// We subtract 0.5 to allow negative movement (range -0.5 to 0.5)
		dx := r.Float64() - 0.5
		dy := r.Float64() - 0.5
		dz := r.Float64() - 0.5

		x += dx
		y += dy
		z += dz
	}

	return []float64{x, y, z}, nil
}
