package internal

/*
this is where we actually read the object imformation. it pulls object and 
object_file together to actually do the stuff you think those files should do
*/
type ObjectDatabase struct {
}

func NewObjectDatabase() *ObjectDatabase{
	return &ObjectDatabase{}
}
