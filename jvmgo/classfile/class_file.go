package classfile


type ClassFile struct {
	Magic uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool ConstantPool
	AccessFlags uint16
	ThisClass uint16
	SuperClass uint16
	Interfaces []uint16
	Fields []*MemberInfo
	Methods []*MemberInfo
	//Attributes []AtrributeInfo
}
