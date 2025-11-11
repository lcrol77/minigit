package internal

import (
	"io/fs"
	"log"
	"os"
)

func stageAllFiles() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fileSys := os.DirFS(wd)
	fs.WalkDir(fileSys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() && d.Name() == ".minigit" {
			return fs.SkipDir
		}
		WriteObjectFile(path, Blob)
		return nil
	})

	return nil
}


