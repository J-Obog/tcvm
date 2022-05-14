package slf

import "bytes"

func (p *Program) Decode(buf *bytes.Buffer) {
	p.decodeHeader(buf)
	p.decodeStrTab(buf)
	p.decodeSymTab(buf)
	p.decodeRelTab(buf)
	p.decodeSegs(buf)
}

func (p *Program) decodeHeader(buf *bytes.Buffer) {
	p.EntryPoint = ReadU32(buf)
	p.StartAddress = ReadU32(buf)
	p.StrTabSize = ReadU32(buf)
	p.SymTabSize = ReadU32(buf)
	p.RelTabSize = ReadU32(buf)
	p.DataSegSize = ReadU32(buf)
	p.CodeSegSize = ReadU32(buf)
}

func (p *Program) decodeStrTab(buf *bytes.Buffer) {
	for i := uint32(0); i < p.StrTabSize; i++ {
		strLen := ReadU32(buf)
		str := ReadStr(buf, strLen)
		p.StrTab = append(p.StrTab, str)
	}
}

func (p *Program) decodeSymTab(buf *bytes.Buffer) {
	for i := uint32(0); i < p.SymTabSize; i++ {
		offSet := ReadU32(buf)
		strIdx := ReadU32(buf)
		flags := ReadU8(buf)
		p.SymTab[p.StrTab[strIdx]] = &Symbol{Offset: offSet, StrTabIndex: strIdx, Flags: byte(flags)}
	}
}

func (p *Program) decodeRelTab(buf *bytes.Buffer) {
	for i := uint32(0); i < p.RelTabSize; i++ {
		offSet := ReadU32(buf)
		strIdx := ReadU32(buf)
		p.RelTab = append(p.RelTab, &Target{Offset: offSet, StrTabIndex: strIdx})
	}
}

func (p *Program) decodeSegs(buf *bytes.Buffer) {
	p.DataSeg = buf.Next(int(p.DataSegSize))
	p.CodeSeg = buf.Next(int(p.CodeSegSize))
}