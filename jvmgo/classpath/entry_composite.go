package classpath

type CompositeEntry []Entry

func newCompositeEntry(path string) *CompositeEntry {
	//compositeEntry := []Entry{}
	return nil  //todo
}

func (self *CompositeEntry) readClass(className string) ([]byte, DirEntry, error) {
	return nil, DirEntry{}, nil
}

func (self *CompositeEntry) String() string {
	return ""
}
