package internal

/*
Note: conceptually how this and object_file differ is this file is responsible
for maintaining objects in memory -- or pieces of memory that we can actually
work with. Object_file is the adapter that interacts with the file system to
actually load that information.
*/
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
	Data     []byte
}

type Oid struct {
	Id []byte
}

type ObjectStore struct {
	objects map[string]*Object
}

// this probably calls ParseObjectFile
// in actual git it will look to see if we have already parsed
// the object yet
func ParseObject(oid Oid) *Object {
	panic("unimplemented")
}

func ObjectTypeFromString(s string) (ObjectTypes, bool) {
	for k, v := range ObjectTypesMap {
		if v == s {
			return k, true
		}
	}
	return 0, false
}
