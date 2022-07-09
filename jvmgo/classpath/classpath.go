package classpath

import "path/filepath"

type ClassPath struct {
	BootClassPath Entry
	ExtClassPath  Entry
	UserClassPath Entry


}

func (self ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.BootClassPath = newWildcardEntry(jreLibPath)

}

func getJreDir(jreOption string) string {

	return ""
}

func (self ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == ""{
		cpOption = "."
	}
	//self.UserClassPath = newEntry(cpOption)
}

func (self ClassPath) ReadClass(name string) ([]byte, Entry,error) {

	return nil,nil,nil
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}
