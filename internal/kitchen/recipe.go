package kitchen

import "context"

// Recipe is a focused media-processing workflow.
type Recipe interface {
	Name() string
	Description() string
	Cook(ctx context.Context, pantry *Pantry) error
}
