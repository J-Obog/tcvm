package exe

type Target struct {
	Offset      uint32
	StrTabIndex uint32
}

type Symbol struct {
	Offset   uint32
	IsExtern bool
	IsData   bool
}

type Header struct {
	EntryPoint   uint32
	StartAddress uint32
	StrTabLen    uint32
	SymTabLen    uint32
	RelTabStart  uint32
	DataSegLen   uint32
	CodeSegLen   uint32
}

type Program struct {
	Header
	CodeSeg []byte
	DataSeg []byte
	StrTab  []string
	SymTab  map[string]*Symbol
	RelTab  []*Target
}
