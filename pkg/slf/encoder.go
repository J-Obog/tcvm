package slf

import "bytes"

func (p *Program) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	p.encodeHeader(buf)
	p.encodeStrTab(buf)
	p.encodeSymTab(buf)
	p.encodeRelTab(buf)
	p.encodeSegs(buf)
	return buf
}

func (p *Program) encodeHeader(buf *bytes.Buffer) {
	p.StrTabSize = uint32(len(p.StrTab))
	p.SymTabSize = uint32(len(p.SymTab))
	p.RelTabSize = uint32(len(p.RelTab))
	p.DataSegSize = uint32(len(p.DataSeg))
	p.CodeSegSize = uint32(len(p.CodeSeg))

	WriteU32(buf, p.EntryPoint)
	WriteU32(buf, p.StartAddress)
	WriteU32(buf, p.StrTabSize)
	WriteU32(buf, p.SymTabSize)
	WriteU32(buf, p.RelTabSize)
	WriteU32(buf, p.DataSegSize)
	WriteU32(buf, p.CodeSegSize)
}

func (p *Program) encodeStrTab(buf *bytes.Buffer) {
	for _, s  := range p.StrTab {
		WriteU32(buf, uint32(len(s)))
		WriteStr(buf, s)
	}
}

func (p *Program) encodeSymTab(buf *bytes.Buffer) {
	for _, s := range p.SymTab {
		WriteU32(buf, s.Offset)
		WriteU32(buf, s.StrTabIndex)
		WriteU8(buf, s.Flags)
	}
}

func (p *Program) encodeRelTab(buf *bytes.Buffer) {
	for _, t := range p.RelTab {
		WriteU32(buf, t)
	}
}

func (p *Program) encodeSegs(buf *bytes.Buffer) {
	buf.Write(p.DataSeg)
	buf.Write(p.CodeSeg)
}