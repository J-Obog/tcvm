package exe

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
	SymTab  SymbolTable
	RelTab  RelocTable
}
