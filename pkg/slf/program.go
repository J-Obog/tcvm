package slf

//SLF = Simple Link Format

type Target struct {
	Offset      uint32
	StrTabIndex uint32
}

type Symbol struct {
	Offset   uint32
	StrTabIndex uint32
	Flags byte //[EXTERN | DATA] 
}

type Header struct {
	EntryPoint   uint32
	StartAddress uint32
	StrTabLen    uint32
	SymTabLen    uint32
	RelTabLen    uint32
	DataSegLen   uint32
	CodeSegLen   uint32
}

const (
	HEADER_START uint32 = 0
	HEADER_LEN uint32 = 28
)

type Program struct {
	Header
	CodeSeg []byte
	DataSeg []byte
	StrTab  []string
	SymTab  map[string]*Symbol
	RelTab  []*Target
}


