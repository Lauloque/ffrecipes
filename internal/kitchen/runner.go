package kitchen

import "context"

// Runner executes an external processing command.
type Runner interface {
	Run(ctx context.Context, args ...string) error
}
