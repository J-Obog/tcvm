package asm

type Assembler struct {
	symtab     map[string]uint32
	codeSeg    []byte
	dataSeg    []byte
	dataRelocs []string
	codeRelocs map[string]uint32
	program    []Statement
}

func NewAssembler(p *Parser) *Assembler {
	a := &Assembler{}
	stmt := p.NextStatement()

	for stmt != nil {
		a.program = append(a.program, stmt)
		stmt = p.NextStatement()
	}

	return a
}

func (a *Assembler) encodeHeader() []byte {
	return nil
}

func (a *Assembler) Assemble(out string) {

}