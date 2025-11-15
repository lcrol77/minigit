package internal

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func ParseObjectFile(oid Oid) (objType ObjectTypes, content []byte, err error) {
	return 1, nil, nil
}

func ReadObjectFile(oid Oid) ([]byte, error) {
	hash := fmt.Sprintf("%x", oid.Id)
	data, err := os.ReadFile(hash) // TODO: eventually replace this with os.Open
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil {
		panic(err)
	}
	return out.Bytes(), nil
}

func WriteObjectFile(filepath string, objType ObjectTypes) error {
	data, err := os.ReadFile(filepath) // TODO: eventually replace this with os.Open
	if err != nil {
		return err
	}
	store := CreateDataStore(data, objType)
	fileHash := ComputeHash(store)
	compressed := compressData(store)
	err = os.WriteFile(fmt.Sprintf(".minigit/objects/%x", fileHash), compressed, 0644)
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

func CreateDataStore(data []byte, objType ObjectTypes) []byte {
	header := generateHeader(data, objType)
	store := append([]byte(header), data...)
	return store
}

func generateHeader(data []byte, objType ObjectTypes) string {
	return fmt.Sprintf("%s %d\x00", ObjectTypesMap[objType], len(data))
}
