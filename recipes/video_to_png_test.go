package recipes

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
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
	dir := t.TempDir()
	input := filepath.Join(dir, "movie.mp4")
	writeTestFile(t, input)

	runner := &recordingRunner{}
	pantry := &kitchen.Pantry{
		Inputs: []string{input},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (VideoToPNG{}).Cook(context.Background(), pantry); err != nil {
		t.Fatalf("Cook returned error: %v", err)
	}

	args := singleCall(t, runner)
	wantOutput := filepath.Join(dir, "movie_%04d.png")
	if got := args[len(args)-1]; got != wantOutput {
		t.Errorf("output pattern = %q, want %q", got, wantOutput)
	}

	if !hasArgPair(args, "-i", input) {
		t.Errorf("ffmpeg args do not contain input %q: %v", input, args)
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

func TestVideoToPNGCookRejectsMissingInput(t *testing.T) {
	input := filepath.Join(t.TempDir(), "missing.mp4")
	runner := &recordingRunner{}

	err := (VideoToPNG{}).Cook(context.Background(), &kitchen.Pantry{
		Inputs: []string{input},
		Log:    silentLogger{},
		FFmpeg: runner,
	})
	if err == nil || !strings.Contains(err.Error(), "input file") {
		t.Fatalf("Cook error = %v, want missing input error", err)
	}
	if len(runner.calls) != 0 {
		t.Errorf("ffmpeg was called %d times, want 0", len(runner.calls))
	}
}

func TestVideoToPNGCookPropagatesRunnerError(t *testing.T) {
	input := filepath.Join(t.TempDir(), "movie.mp4")
	writeTestFile(t, input)

	runner := &recordingRunner{err: errors.New("boom")}
	pantry := &kitchen.Pantry{
		Inputs: []string{input},
		Log:    silentLogger{},
		FFmpeg: runner,
	}

	if err := (VideoToPNG{}).Cook(context.Background(), pantry); err == nil {
		t.Fatal("Cook returned nil error after runner failure")
	}
}
