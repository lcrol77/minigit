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

