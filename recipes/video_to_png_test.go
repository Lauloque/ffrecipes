package recipes

import (
	"context"
	"errors"
	"testing"

	"github.com/Lauloque/ffrecipes/internal/kitchen"
)

type recordingRunner struct {
	calls [][]string
	err   error
}

func (r *recordingRunner) Run(_ context.Context, args ...string) error {
	r.calls = append(r.calls, args)
	return r.err
}

type silentLogger struct{}

func (silentLogger) Info(string)    {}
func (silentLogger) Success(string) {}
func (silentLogger) Warning(string) {}
func (silentLogger) Error(string)   {}
func (silentLogger) Debug(string)   {}

func TestVideoToPNGCook(t *testing.T) {
	runner := &recordingRunner{}
	pantry := &kitchen.Pantry{
		Inputs: []string{"movie.mp4"},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (VideoToPNG{}).Cook(context.Background(), pantry); err != nil {
		t.Fatalf("Cook returned error: %v", err)
	}

	if len(runner.calls) != 1 {
		t.Fatalf("got %d ffmpeg calls, want 1", len(runner.calls))
	}

	args := runner.calls[0]
	if got, want := args[len(args)-1], "movie_%04d.png"; got != want {
		t.Errorf("output pattern = %q, want %q", got, want)
	}

	foundInput := false
	for i, arg := range args {
		if arg == "-i" && i+1 < len(args) && args[i+1] == "movie.mp4" {
			foundInput = true
			break
		}
	}
	if !foundInput {
		t.Errorf("ffmpeg args do not contain input movie.mp4: %v", args)
	}
}

func TestVideoToPNGCookRequiresInput(t *testing.T) {
	err := (VideoToPNG{}).Cook(context.Background(), &kitchen.Pantry{
		Log:    silentLogger{},
		FFmpeg: &recordingRunner{},
	})
	if err == nil {
		t.Fatal("Cook returned nil error without inputs")
	}
}

func TestVideoToPNGCookPropagatesRunnerError(t *testing.T) {
	runner := &recordingRunner{err: errors.New("boom")}
	pantry := &kitchen.Pantry{
		Inputs: []string{"movie.mp4"},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (VideoToPNG{}).Cook(context.Background(), pantry); err == nil {
		t.Fatal("Cook returned nil error after runner failure")
	}
}
