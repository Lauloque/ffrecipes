package dishes

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Lauloque/ffrecipes/internal/kitchen"
)

// VideoToPNG extracts PNG frames from video files.
type VideoToPNG struct{}

func (VideoToPNG) Name() string {
	return "video-to-png"
}

func (VideoToPNG) Description() string {
	return "Convert video files to PNG image sequences."
}

func (VideoToPNG) Cook(ctx context.Context, pantry *kitchen.Pantry) error {
	if len(pantry.Inputs) == 0 {
		return fmt.Errorf("no input files provided")
	}
	if pantry.FFmpeg == nil {
		return fmt.Errorf("ffmpeg runner is not configured")
	}

	for _, input := range pantry.Inputs {
		output := strings.TrimSuffix(input, filepath.Ext(input)) + "_%04d.png"
		pantry.Log.Info(fmt.Sprintf("Extracting PNG frames from %q...", input))

		args := []string{
			"-hide_banner",
			"-loglevel", "error",
			"-nostdin",
			"-y",
			"-i", input,
			"-fps_mode", "passthrough",
			output,
		}

		if err := pantry.FFmpeg.Run(ctx, args...); err != nil {
			return fmt.Errorf("convert %q: %w", input, err)
		}

		pantry.Log.Success(fmt.Sprintf("Wrote %q.", output))
	}

	return nil
}

func init() {
	kitchen.Register(VideoToPNG{})
}
