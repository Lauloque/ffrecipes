package files

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRequireAcceptsFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "movie.mp4")
	if err := os.WriteFile(path, []byte("test"), 0o600); err != nil {
		t.Fatalf("write test file: %v", err)
	}

	if err := Require(path, "video file"); err != nil {
		t.Fatalf("Require returned error: %v", err)
	}
}

func TestRequireRejectsMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.mp4")

	err := Require(path, "video file")
	if err == nil {
		t.Fatal("Require returned nil error for missing file")
	}
	if !strings.Contains(err.Error(), "video file") {
		t.Errorf("error = %q, want it to mention the label", err)
	}
}

func TestRequireRejectsDirectory(t *testing.T) {
	path := t.TempDir()

	err := Require(path, "video file")
	if err == nil || !strings.Contains(err.Error(), "is a directory") {
		t.Fatalf("Require error = %v, want directory error", err)
	}
}
