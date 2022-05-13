package exe

type Target struct {
	Offset uint32
	Index  uint32
}

type RelocTable []*Target

type Symbol struct {
	Label    string
	Offset   uint32
	IsExtern bool
	IsData   bool
}

type SymbolTable []*Symbol

func (st *SymbolTable) Get(lbl string) *Symbol {
	for _, sym := range *st {
		if (sym != nil) && (sym.Label == lbl) {
			return sym
		}
	}
	return nil
}

func (st *SymbolTable) Add(sym *Symbol) bool {
	if st.Get(sym.Label) != nil {
		return false
	}

	*st = append(*st, sym)
	return true
}

func (st *SymbolTable) Remove(lbl string) bool {
	sym := st.Get(lbl)

	if sym == nil {
		return false
	}

	sym = nil
	return true
}