package asm

type Assembler struct {
	size    uint32
	symtab  map[string]uint32
	program []Statement
}

func NewAssembler(p *Parser) *Assembler {
	a := &Assembler{}
	stmt := p.NextStatement()
	loc := uint32(0)

	for stmt != nil {
		switch stmt := stmt.(type) {
		case *Label:
			if _, ok := a.symtab[stmt.Name]; ok { //checking if label has been declared already
				panic("Redefinition of symbol")
			}
			a.symtab[stmt.Name] = loc

		case *Data, *Instruction:
			a.program = append(a.program, stmt)
		}

		loc += uint32(stmt.TotalSize())
		stmt = p.NextStatement()
	}

	a.size = loc
	return a
}

func (a *Assembler) encodeHeader() []byte {
	return nil
}

func (a *Assembler) Assemble(out string) {

}