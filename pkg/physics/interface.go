package physics

// Simulation defines the interface for any physics model we want to run.
// This makes our worker flexible: we can swap Random Walk for Ising Model later.
type Simulation interface {
	// Run executes the simulation and returns the result coordinates.
	// seed: Ensures reproducibility.
	// iterations: How long to run the simulation.
	Run(seed int64, iterations int64) ([]float64, error)
}
