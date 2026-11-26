package recipes

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Lauloque/ffrecipes/internal/kitchen"
)

func TestAddCoverToVideoUsesCoverNextToVideo(t *testing.T) {
	t.Chdir(t.TempDir())

	videoDir := t.TempDir()
	video := filepath.Join(videoDir, "movie.mp4")
	cover := filepath.Join(videoDir, "cover.jpg")
	writeTestFile(t, video)
	writeTestFile(t, cover)

	runner := &recordingRunner{}
	pantry := &kitchen.Pantry{
		Inputs: []string{video},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (AddCoverToVideo{}).Cook(context.Background(), pantry); err != nil {
		t.Fatalf("Cook returned error: %v", err)
	}

	args := singleCall(t, runner)
	if !hasArgPair(args, "-i", video) {
		t.Errorf("ffmpeg args do not contain input %q: %v", video, args)
	}
	if !hasArgPair(args, "-i", cover) {
		t.Errorf("ffmpeg args do not contain default cover %q: %v", cover, args)
	}
	if !hasArgPair(args, "-c", "copy") {
		t.Errorf("ffmpeg args do not copy the original streams: %v", args)
	}
	if !hasArgPair(args, "-c:v:1", "png") {
		t.Errorf("ffmpeg args do not set the attached picture codec: %v", args)
	}
	if !hasArgPair(args, "-disposition:v:1", "attached_pic") {
		t.Errorf("ffmpeg args do not mark the cover as attached: %v", args)
	}

	wantOutput := filepath.Join(videoDir, "movie.cover.mp4")
	if got := args[len(args)-1]; got != wantOutput {
		t.Errorf("output = %q, want %q", got, wantOutput)
	}
}

func TestAddCoverToVideoUsesExplicitCoverMIME(t *testing.T) {
	videoDir := t.TempDir()
	coverDir := t.TempDir()
	video := filepath.Join(videoDir, "movie.mkv")
	cover := filepath.Join(coverDir, "art.png")
	writeTestFile(t, video)
	writeTestFile(t, cover)

	runner := &recordingRunner{}
	pantry := &kitchen.Pantry{
		Inputs: []string{video, cover},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (AddCoverToVideo{}).Cook(context.Background(), pantry); err != nil {
		t.Fatalf("Cook returned error: %v", err)
	}

	args := singleCall(t, runner)
	if !hasArgPair(args, "-attach", cover) {
		t.Errorf("ffmpeg args do not attach cover %q: %v", cover, args)
	}
	if !hasArgPair(args, "-metadata:s:t", "mimetype=image/png") {
		t.Errorf("ffmpeg args do not set the PNG MIME type: %v", args)
	}

	wantOutput := filepath.Join(videoDir, "movie.cover.mkv")
	if got := args[len(args)-1]; got != wantOutput {
		t.Errorf("output = %q, want %q", got, wantOutput)
	}
}

func TestAddCoverToVideoRejectsMissingFiles(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T) []string
		wantErr string
	}{
		{
			name: "video",
			setup: func(t *testing.T) []string {
				return []string{filepath.Join(t.TempDir(), "missing.mp4")}
			},
			wantErr: "video file",
		},
		{
			name: "default cover",
			setup: func(t *testing.T) []string {
				video := filepath.Join(t.TempDir(), "movie.mp4")
				writeTestFile(t, video)
				return []string{video}
			},
			wantErr: "cover file",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runner := &recordingRunner{}
			pantry := &kitchen.Pantry{
				Inputs: test.setup(t),
				Log:    silentLogger{},
				FFmpeg: runner,
			}

			err := (AddCoverToVideo{}).Cook(context.Background(), pantry)
			if err == nil {
				t.Fatal("Cook returned nil error")
			}
			if !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, test.wantErr)
			}
			if len(runner.calls) != 0 {
				t.Errorf("ffmpeg was called %d times, want 0", len(runner.calls))
			}
		})
	}
}

func TestAddCoverToVideoRejectsUnsupportedFormats(t *testing.T) {
	dir := t.TempDir()
	video := filepath.Join(dir, "movie.mp4")
	cover := filepath.Join(dir, "cover.gif")
	writeTestFile(t, video)
	writeTestFile(t, cover)

	err := (AddCoverToVideo{}).Cook(context.Background(), &kitchen.Pantry{
		Inputs: []string{video, cover},
		Log:    silentLogger{},
		FFmpeg: &recordingRunner{},
	})
	if err == nil || !strings.Contains(err.Error(), "unsupported cover format") {
		t.Fatalf("Cook error = %v, want unsupported cover format", err)
	}
}

func TestAddCoverToVideoRejectsExtraInputs(t *testing.T) {
	err := (AddCoverToVideo{}).Cook(context.Background(), &kitchen.Pantry{
		Inputs: []string{"one.mp4", "cover.jpg", "extra"},
		Log:    silentLogger{},
		FFmpeg: &recordingRunner{},
	})
	if err == nil || !strings.Contains(err.Error(), "got 3 inputs") {
		t.Fatalf("Cook error = %v, want extra input error", err)
	}
}

func writeTestFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("test"), 0o600); err != nil {
		t.Fatalf("write test file %q: %v", path, err)
	}
}

func singleCall(t *testing.T, runner *recordingRunner) []string {
	t.Helper()
	if len(runner.calls) != 1 {
		t.Fatalf("got %d ffmpeg calls, want 1", len(runner.calls))
	}
	return runner.calls[0]
}

func hasArgPair(args []string, flag, value string) bool {
	for i, arg := range args {
		if arg == flag && i+1 < len(args) && args[i+1] == value {
			return true
		}
	}
	return false
}
