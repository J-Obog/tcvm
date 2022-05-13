package exe

type Symbol struct {
	Label    string
	Offset   uint32
	IsExtern bool
	IsData   bool
}

type Target struct {
	Offset uint32
	Index  uint32
}

type RelocTable []*Target

type SymbolTable []*Symbol
