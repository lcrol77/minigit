package internal_test

import (
	"minigit/cmd/minigit_cli/internal"
	"minigit/cmd/minigit_cli/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInitializeRepo(t *testing.T) {
	_, cleanup := setupEmptyRepoDir(t) // setup repo without initializing it
	defer cleanup()

	// Now call Init() to create the repo structure
	if err := internal.Init(); err != nil {
		t.Fatalf("Error initializing repo: %v", err)
	}

	// List of files and directories we expect after initialization
	expectedPaths := []string{
		".minigit/objects",
		".minigit/commits",
		".minigit/refs",
		".minigit/HEAD",
	}

	for _, path := range expectedPaths {
		if exists, err := testutil.Exists(path); err != nil {
			t.Fatalf("Error checking %s: %v", path, err)
		} else if !exists {
			t.Fatalf("Expected %s to exist, but it does not", path)
		}
	}
}

// setupEmptyRepoDir creates a temporary repo directory without calling Init.
// It returns the repo path and a cleanup function.
func setupEmptyRepoDir(t *testing.T) (string, func()) {
	tmpDir := filepath.Join(os.TempDir(), "repo-init-test")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatalf("could not create test repo dir: %v", err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("could not chdir to repo: %v", err)
	}

	cleanup := func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Logf("cleanup failed: %v", err)
		}
	}
	return tmpDir, cleanup
}
