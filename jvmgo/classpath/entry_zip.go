package classpath

import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry{
	absDir,err := filepath.Abs(path)
	if err !=nil{
		panic(err)
	}
	return &ZipEntry{absDir}
}

func(self *ZipEntry) String() string{
	return self.absPath
}