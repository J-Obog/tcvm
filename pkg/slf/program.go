package slf

//SLF = Simple Link Format

type Symbol struct {
	Offset   uint32
	StrTabIndex uint32
	Flags uint8 //[EXTERN | DATA] 
}

const ( //symbol flag mapping
	S_ISEXTERN uint8 = 0x1 
	S_ISDATA uint8 = 0x2
)

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
	RelTab  []uint32
}


