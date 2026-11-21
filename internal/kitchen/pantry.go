package kitchen

import "github.com/Lauloque/ffrecipes/internal/ui"

type Pantry struct {
	Inputs []string

	FFmpegPath string

	Log ui.Logger

	Options Options
}
