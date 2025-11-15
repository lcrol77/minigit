package testutil

import (
	"errors"
	"io/fs"
	"minigit/cmd/minigit_cli/internal"
	"os"
	"path/filepath"
)

// returns an path to an initialized repo
func InitializeTestingRepo() (string, error) {
	tmpDir := filepath.Join(os.TempDir(), "repo")
	err := os.Mkdir(tmpDir, 0755)
	if err != nil {
		return "", err
	}
	fp := filepath.Join(tmpDir, "test.txt")
	err = os.WriteFile(fp, []byte("test\n"), 0644)
	if err != nil {
		return "", err
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return "", err
	}
	err = internal.Init()
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}

func CleanUpRepo() error {
	tmpDir := os.TempDir()
	err := os.RemoveAll(filepath.Join(tmpDir, "repo"))
	return err
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}
