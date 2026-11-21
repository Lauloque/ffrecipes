package kitchen

type Pantry struct {
	Inputs []string

	FFmpegPath string

	Log Logger

	Options Options
}
