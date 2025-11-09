package internal

type ObjectTypes int
const (
	Blob ObjectTypes = iota
	Tree
)

var ObjectTypesMap = map[ObjectTypes]string {
	Blob: "blob",
}

type Object struct {
	ObjectID Oid
	Type ObjectTypes
}

type Oid struct {
	Id *[]byte
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

