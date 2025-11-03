package repo

import (
	"fmt"
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
	hashedFile := computeFileHash(&data)
	compressed := compressData(data)
	err = os.WriteFile(fmt.Sprintf(".minigit/objects/%x", hashedFile), compressed, 0644)
	if err != nil {
		return err
	}
	return nil
}

