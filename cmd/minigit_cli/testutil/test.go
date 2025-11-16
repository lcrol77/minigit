package testutil

import (
	"errors"
	"io/fs"
	"minigit/cmd/minigit_cli/internal"
	"os"
	"path/filepath"
	"testing"
)

// InitializeTestingRepo creates an isolated repo in a temp directory and returns its path.
func InitializeTestingRepo() (string, error) {
	tmpDir, err := os.MkdirTemp("", "repo-*")
	if err != nil {
		return "", err
	}

	// Create a simple test file
	if err := os.WriteFile(filepath.Join(tmpDir, "test.txt"), []byte("test\n"), 0644); err != nil {
		return "", err
	}

	// Initialize minigit inside the repo
	if err := runInitInPath(tmpDir); err != nil {
		return "", err
	}

	return tmpDir, nil
}

// Initializes minigit in the specified path
func runInitInPath(path string) error {
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)

	if err := os.Chdir(path); err != nil {
		return err
	}

	return internal.Init()
}

// Removes the entire repo; used as cleanup
func CleanUpRepo(repoPath string) {
	_ = os.RemoveAll(repoPath) // ignore cleanup errors for now
}

// Checks whether a file exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return err == nil, err
}

// SetupAndChdir calls InitializeTestingRepo and chdirs into it, returning repo path and cleanup func.
func SetupAndChdir(t *testing.T) (string, func()) {
	repo, err := InitializeTestingRepo()
	if err != nil {
		t.Fatalf("InitializeTestingRepo failed: %v", err)
	}

	if err := os.Chdir(repo); err != nil {
		CleanUpRepo(repo)
		t.Fatalf("Chdir failed: %v", err)
	}

	cleanup := func() { CleanUpRepo(repo) }
	return repo, cleanup
}
