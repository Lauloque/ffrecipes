package kitchen

// Pantry holds the inputs and services available to a recipe.
type Pantry struct {
	Inputs []string

	Log Logger

	FFmpeg Runner
}
