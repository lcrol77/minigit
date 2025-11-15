package internal

type ObjectTypes int

const (
	Blob ObjectTypes = iota
	Tree
)

var ObjectTypesMap = map[ObjectTypes]string{
	Blob: "blob",
	Tree: "tree",
}

type Object struct {
	ObjectID Oid
	Type     ObjectTypes
}

type Oid struct {
	Id []byte
}

func ObjectTypeFromString(s string) (ObjectTypes, bool) {
	for k, v := range ObjectTypesMap {
		if v == s {
			return k, true
		}
	}
	return 0, false
}

func LookUpObject(oid Oid) *Object {
	panic("unimplemented")
}

func ParseObject(oid Oid) *Object {
	panic("unimplemented")
}

func CreateObject(oid Oid) *[]byte {
	panic("unimplemented")
}
