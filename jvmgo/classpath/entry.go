package classpath

import (
	"os"
)

const pathListSeparator = string(os.PathSeparator)

type  Entry interface {
	readClass(className string) ([]byte, Entry,error)
	String() string
}


//todo check
//func newEntry(path string) *Entry {
	//if strings.Contains(path,pathListSeparator){
	//	return newCompositeEntry(path)
	//}
	//if strings.Contains(path,"*"){
		//return newWildcardEntry(path)
	//}
	//if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") ||
	//	strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP"){
	//	return newZipEntry(path)
	//}
	//return newDirEntry(path)
//}