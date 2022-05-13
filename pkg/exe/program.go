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
	SymTabStart  uint32
	RelTabStart  uint32
	DataSegStart uint32
	CodeSegStart uint32
}

type Program struct {
	Header
	CodeSeg []byte
	DataSeg []byte
	StrTab  []string
	SymTab  map[string]*Symbol
	RelTab  []*Target
}
