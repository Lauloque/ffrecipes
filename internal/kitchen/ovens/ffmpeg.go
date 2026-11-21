package ovens

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// FFmpeg runs the FFmpeg executable.
type FFmpeg struct {
	Executable string
}

// NewFFmpeg returns an FFmpeg runner for executable.
func NewFFmpeg(executable string) *FFmpeg {
	return &FFmpeg{Executable: executable}
}

// Run executes FFmpeg with args and waits for it to finish.
func (f *FFmpeg) Run(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, f.Executable, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}

		message := strings.TrimSpace(string(output))
		if message == "" {
			return fmt.Errorf("ffmpeg: %w", err)
		}
		return fmt.Errorf("ffmpeg: %w: %s", err, message)
	}

	return nil
}
