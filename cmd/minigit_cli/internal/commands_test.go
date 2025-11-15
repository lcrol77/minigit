package internal_test

import (
	"minigit/cmd/minigit_cli/internal"
	"minigit/cmd/minigit_cli/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInitializeRepo(t *testing.T) {
	testCases := []string{
		".minigit/objects",
		".minigit/commits",
		".minigit/refs",
		".minigit/HEAD",
	}
	tmpDir := os.TempDir()
	err := os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("could not make tmpdir: %v", err)
	}
	err = os.Mkdir("repo", 0755)
	if err != nil {
		t.Fatalf("could not make tmpdir: %v", err)
	}
	defer testutil.CleanUpRepo()

	tmpDir = filepath.Join(tmpDir, "repo")
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Could change to %s: %v", tmpDir,err)
	}

	err = internal.Init()
	if err != nil {
		t.Fatalf("Error initializing repo: %v", err)
	}
	for _, path := range testCases {
		exist, err := testutil.Exists(path)
		if err != nil {
			t.Fatalf("repo malformed with error: %v", err)
		}
		if !exist {
			t.Fatalf("repo malformed: expected %s to exist", path)
		}
	}
}
