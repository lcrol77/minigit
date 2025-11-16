package internal

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ParseObjectFile(oid Oid) (objType ObjectTypes, content []byte, err error) {
	raw, err := ReadObjectFile(oid)
	if err != nil {
		return objType, nil, err
	}
	spaceIdx := bytes.IndexByte(raw, ' ')
	if spaceIdx < 0 {
		return objType, nil, fmt.Errorf("invalid object header: no space delimeter")
	}
	objTypeString := string(raw[:spaceIdx])
	objType, found := ObjectTypeFromString(objTypeString)
	if !found {
		return objType, nil, fmt.Errorf("invalid object header: no type found")
	}

	nullIdx := bytes.IndexByte(raw, 0)
	if nullIdx < 0 {
		return objType, nil, fmt.Errorf("invalid object header: no null byte found")
	}

	// Content starts just after the null byte
	content = raw[nullIdx+1:]
	return objType, content, nil
}

func getObjectFilePath(h string) string {
	return filepath.Join(".minigit/objects/", h)
}

func ReadObjectFile(oid Oid) ([]byte, error) {
	path := getObjectFilePath(fmt.Sprintf("%x", oid.Id))
	data, err := os.ReadFile(path) // TODO: eventually replace this with os.Open
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

func WriteObjectFile(filepath string, objType ObjectTypes) (Oid, error) {
	data, err := os.ReadFile(filepath) // TODO: eventually replace this with os.Open
	if err != nil {
		return Oid{}, err
	}
	store := CreateDataStore(data, objType)
	fileHash := ComputeHash(store)
	compressed := compressData(store)
	
	oid := Oid{Id: fileHash}

	err = os.WriteFile(fmt.Sprintf(".minigit/objects/%x", fileHash), compressed, 0644)
	if err != nil {
		return Oid{}, err
	}
	return oid, nil
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
