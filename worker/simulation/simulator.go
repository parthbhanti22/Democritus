package simulation

// Simulator defines the interface for any physics model.
// This allows us to easily swap the "Physics Engine" (Strategy Pattern).
type Simulator interface {
	// Run executes the simulation.
	// seed: Ensures reproducibility (CERN requirement).
	// iterations: Number of steps.
	// Returns: A slice of coordinates [x, y, z, ...].
	Run(seed int64, iterations int64) ([]float64, error)
}
