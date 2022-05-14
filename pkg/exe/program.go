package exe

import "encoding/binary"

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

const HEADER_LEN uint32 = 28
const HEADER_START uint32 = 0

type Program struct {
	Header
	CodeSeg []byte
	DataSeg []byte
	StrTab  []string
	SymTab  map[string]*Symbol
	RelTab  []*Target
}

func stou32(bslice []byte, littleEndian bool) uint32 {
	if littleEndian {
		return binary.LittleEndian.Uint32(bslice)
	} else {
		return binary.BigEndian.Uint32(bslice)
	}
}

func (p *Program) Decode(o []byte) {
	p.decodeHeader(o)
	p.decodeStrTab(o)	
	p.decodeSymTab(o)
	p.decodeRelTab(o)
	p.decodeSegs(o)
}

func (p *Program) decodeHeader(o []byte) {
	pos := HEADER_START
	
	p.EntryPoint = stou32(o[pos: pos + 4], true)
	pos += 4
	p.StartAddress = stou32(o[pos: pos + 4], true)
	pos += 4
	p.StrTabLen = stou32(o[pos: pos + 4], true)
	pos += 4
	p.SymTabLen = stou32(o[pos: pos + 4], true)
	pos += 4
	p.RelTabLen = stou32(o[pos: pos + 4], true)
	pos += 4
	p.DataSegLen = stou32(o[pos: pos + 4], true)
	pos += 4
	p.CodeSegLen = stou32(o[pos: pos + 4], true)	
	pos += 4
}

func (p *Program) decodeStrTab(o []byte) {
	pos := HEADER_LEN

	for pos != p.StrTabLen {
		var buf []byte
		strLen := stou32(o[pos:pos + 4], true)
		pos += 4
		for pos < (pos+strLen) { 
			buf = append(buf, o[pos])
			pos += 1
		}
		p.StrTab = append(p.StrTab, string(buf))
	}
}

func (p *Program) decodeSymTab (o []byte) {
	pos := HEADER_LEN
	pos += p.StrTabLen

	for pos != p.SymTabLen {
		sym := &Symbol{}
		
		offSet := stou32(o[pos: pos + 4], true)
		sym.Offset = offSet
		pos += 4

		strIdx := stou32(o[pos: pos + 4], true)
		sym.StrTabIndex = strIdx
		pos += 4

		sym.Flags = o[pos]
		pos += 1

		p.SymTab[p.StrTab[strIdx]] = sym
	}
}

func (p *Program) decodeRelTab (o []byte) {
	pos := HEADER_LEN
	pos += p.StrTabLen  
	pos += p.SymTabLen

	for pos != p.RelTabLen {
		target := &Target{}

		offSet := stou32(o[pos: pos + 4], true)		
		target.Offset = offSet
		pos += 4

		strIdx := stou32(o[pos: pos + 4], true)
		target.StrTabIndex = strIdx
		pos += 4

		p.RelTab = append(p.RelTab, target)
	}
}


func (p *Program) decodeSegs(o []byte) {
	pos := HEADER_LEN
	pos += p.StrTabLen  
	pos += p.SymTabLen
	pos += p.RelTabLen

	p.DataSeg = o[pos : pos + p.DataSegLen]
	pos += p.DataSegLen

	p.CodeSeg = o[pos : pos + p.CodeSegLen]
}