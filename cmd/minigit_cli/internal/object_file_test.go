package internal_test

import (
	"fmt"
	"minigit/cmd/minigit_cli/internal"
	"minigit/cmd/minigit_cli/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	repo, err := testutil.InitializeTestingRepo()
	defer testutil.CleanUpRepo()
	if err != nil {
		t.Fatalf("failed to initialize the repo: %v", err)
	}
	err = os.Chdir(repo)
	if err != nil {
		t.Fatalf("failed to initialize the repo: %v", err)
	}
	fp := filepath.Join(repo, "test.txt")
	err = internal.WriteObjectFile(fp, internal.Blob)
	if err != nil {
		t.Fatalf("Failed to write obj file: %v", err)
	}
	store := internal.CreateDataStore([]byte("test\n"), internal.Blob)
	h := internal.ComputeHash(store)
	path := filepath.Join(repo,fmt.Sprintf(".minigit/objects/%x", h))
	e, err := testutil.Exists(path)
	if err != nil {
		t.Fatalf("failed to write with error: %v", err)
	}
	if !e {
		t.Fatalf("failed to write: expected %s to exist", path)
	}
}

func TestReadHeader(t *testing.T) {
	t.Fatalf("unimplemented")
}

func TestReadFileContents(t *testing.T) {
	t.Fatalf("unimplemented")
}
