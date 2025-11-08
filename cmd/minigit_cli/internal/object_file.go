package internal

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)


func ReadObjectFile(fileHash string) (*[]byte, error) {
	fmt.Println(fileHash)
	data, err := os.ReadFile(fileHash) // TODO: eventually replace this with os.Open
	if err != nil {
		return nil, err
	}
	decompressedData := decompressData(data)	
	return &decompressedData, nil
}

func WriteObjectFile(filepath string) error {
	fmt.Println(filepath)
	data, err := os.ReadFile(filepath) // TODO: eventually replace this with os.Open
	if err != nil {
		return err
	}
	hashedFile := computeHash(&data)
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
