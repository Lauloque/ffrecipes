package recipes

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Lauloque/ffrecipes/internal/files"
	"github.com/Lauloque/ffrecipes/internal/kitchen"
)

// AddCoverToVideo adds a cover to MP4 or MKV videos.
type AddCoverToVideo struct{}

func (AddCoverToVideo) Name() string {
	return "add-cover-to-video"
}

func (AddCoverToVideo) Description() string {
	return "Add a cover to an MP4 or MKV video."
}

func (AddCoverToVideo) Cook(ctx context.Context, pantry *kitchen.Pantry) error {
	if len(pantry.Inputs) == 0 {
		return fmt.Errorf("no input files provided")
	}
	if pantry.FFmpeg == nil {
		return fmt.Errorf("ffmpeg runner is not configured")
	}

	video := pantry.Inputs[0]
	var cover string

	switch len(pantry.Inputs) {
	case 1:
		cover = filepath.Join(filepath.Dir(video), "cover.jpg")
	case 2:
		cover = pantry.Inputs[1]
	default:
		return fmt.Errorf("expected one video and an optional cover, got %d inputs", len(pantry.Inputs))
	}

	if err := files.Require(video, "video file"); err != nil {
		return err
	}
	if err := files.Require(cover, "cover file"); err != nil {
		return err
	}

	mime, err := coverMIME(cover)
	if err != nil {
		return err
	}

	rawExt := filepath.Ext(video)
	ext := strings.ToLower(rawExt)
	output := strings.TrimSuffix(video, rawExt) + ".cover" + rawExt

	var args []string

	switch ext {
	case ".mp4":
		args = []string{
			"-hide_banner",
			"-loglevel", "error",
			"-nostdin",
			"-y",
			"-i", video,
			"-i", cover,
			"-map", "0",
			"-map", "1",
			"-c", "copy",
			"-c:v:1", "png",
			"-disposition:v:1", "attached_pic",
			output,
		}
	case ".mkv":
		args = []string{
			"-hide_banner",
			"-loglevel", "error",
			"-nostdin",
			"-y",
			"-i", video,
			"-attach", cover,
			"-metadata:s:t", "mimetype=" + mime,
			"-c", "copy",
			"-map", "0",
			output,
		}
	default:
		return fmt.Errorf("unsupported video format %q: only MP4 and MKV are supported", rawExt)
	}

	pantry.Log.Info(fmt.Sprintf("Adding cover %q to %q...", cover, video))

	if err := pantry.FFmpeg.Run(ctx, args...); err != nil {
		return fmt.Errorf("add cover to %q: %w", video, err)
	}

	pantry.Log.Success(fmt.Sprintf("Wrote %q.", output))

	return nil
}

func coverMIME(path string) (string, error) {
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".jpg", ".jpeg":
		return "image/jpeg", nil
	case ".png":
		return "image/png", nil
	default:
		return "", fmt.Errorf("unsupported cover format %q: use JPG or PNG", ext)
	}
}

func init() {
	kitchen.Register(AddCoverToVideo{})
}
