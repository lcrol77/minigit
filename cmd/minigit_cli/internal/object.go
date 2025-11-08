package internal

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

type ObjectTypes int
const (
	Blob ObjectTypes = iota
	Tree
)

type Object struct {
	ObjectID Oid
	Type ObjectTypes
}

type Oid struct {
	id *[]byte
}

func LookUpObject(oid Oid) *Object{
	panic("unimplemented")
}

func ParseObject(oid Oid) *Object{
	panic("unimplemented")
}

func CreateObject(oid Oid) *[]byte {
	panic("unimplemented")
}

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

func computeHash(data *[]byte) []byte {
	h := sha1.New()
	h.Write(*data)
	sum := h.Sum(nil)
	return sum
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
