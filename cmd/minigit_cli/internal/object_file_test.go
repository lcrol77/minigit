package internal_test

import (
	"bytes"
	"fmt"
	"minigit/cmd/minigit_cli/internal"
	"minigit/cmd/minigit_cli/testutil"
	"path/filepath"
	"testing"
)

func TestWriteObjectFile(t *testing.T) {
	repo, cleanup := testutil.SetupAndChdir(t)
	defer cleanup()

	testFile := filepath.Join(repo, "test.txt")
	oid, err := internal.WriteObjectFile(testFile, internal.Blob)
	if err != nil {
		t.Fatalf("WriteObjectFile failed: %v", err)
	}

	objectPath := filepath.Join(repo, fmt.Sprintf(".minigit/objects/%x", oid.Id))
	exists, err := testutil.Exists(objectPath)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Errorf("Expected object file %s to exist, but it does not", objectPath)
	}
}

func TestParseObjectFile(t *testing.T) {
	repo, cleanup := testutil.SetupAndChdir(t)
	defer cleanup()

	testFile := filepath.Join(repo, "test.txt")
	oid, err := internal.WriteObjectFile(testFile, internal.Blob)
	if err != nil {
		t.Fatalf("WriteObjectFile failed: %v", err)
	}

	objType, content, err := internal.ParseObjectFile(oid)
	if err != nil {
		t.Fatalf("ParseObjectFile failed: %v", err)
	}

	if objType != internal.Blob {
		t.Errorf("Expected type Blob, got %s", internal.ObjectTypesMap[objType])
	}

	expectedContent := []byte("test\n")
	if !bytes.Equal(content, expectedContent) {
		t.Errorf("Expected content %q, got %q", expectedContent, content)
	}
}
