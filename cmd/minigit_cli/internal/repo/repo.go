package repo

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func Init() error {
	err := os.Mkdir(".minigit", 0755)
	if err != nil {
		return err
	}

	os.Mkdir(".minigit/objects", 0755)
	os.Mkdir(".minigit/commits", 0755)
	os.Mkdir(".minigit/refs", 0755)
	os.WriteFile(".minigit/HEAD", []byte("refs/heads/main"), 0644)
	return nil
}

func Add(cmd string) error {
	if cmd == "." {
		return stageAllFiles()
	} else {
		return stageFile(cmd)
	}
}

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
		if d.IsDir() && d.Name() == ".minigit"{
			return fs.SkipDir
		}
		stageFile(path)
		return nil
	})

	return nil
}

func stageFile(filepath string) error {
	fmt.Println(filepath)
	data, err := os.ReadFile(filepath) // TODO: eventually replace this with os.Open
	if err != nil {
		return err
	}
	hashedFile := computeFileHash(&data)
	compressed := compressData(data)
	err = os.WriteFile(fmt.Sprintf(".minigit/objects/%x", hashedFile), compressed, 0644)
	if err != nil {
		return err
	}
	return nil
}

func compressData(data []byte) []byte {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	_, err := zw.Write(data)
	if err != nil {
		panic(err)
	}
	zw.Close()
	return buf.Bytes()
}

func decompressData(compressedData []byte) []byte {
	r, err := zlib.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	decompressed, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return decompressed
}

func computeFileHash(data *[]byte) []byte {
	h := sha1.New()
	h.Write(*data)
	sum := h.Sum(nil)
	fmt.Printf("SHA1: %x\n", sum)
	return sum
}
