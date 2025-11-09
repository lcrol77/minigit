package internal

import (
	"os"
)

func Init() error {
	err := os.Mkdir(".minigit", 0755)
	if err != nil {
		return err
	}

	os.Mkdir(".minigit/objects", 0755)
	//os.Mkdir(".minigit/commits", 0755)
	//os.Mkdir(".minigit/refs", 0755)
	os.WriteFile(".minigit/HEAD", []byte("refs/heads/main"), 0644)
	return nil
}

func Add(cmd string) error {
	if cmd == "." {
		return stageAllFiles()
	} else {
		return WriteObjectFile(cmd, Blob)
	}
}

func Cat(filePath string) error {
	return nil
}
