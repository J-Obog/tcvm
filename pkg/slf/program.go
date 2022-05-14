package slf

//SLF = Simple Link Format

type Target struct {
	Offset      uint32
	StrTabIndex uint32
}

type Symbol struct {
	Offset   uint32
	StrTabIndex uint32
	Flags uint8 //[EXTERN | DATA] 
}

type Header struct {
	EntryPoint   uint32
	StartAddress uint32
	StrTabSize    uint32
	SymTabSize    uint32
	RelTabSize    uint32
	DataSegSize  uint32
	CodeSegSize  uint32
}

type Program struct {
	Header
	CodeSeg []byte
	DataSeg []byte
	StrTab  []string
	SymTab  map[string]*Symbol
	RelTab  []*Target
}


