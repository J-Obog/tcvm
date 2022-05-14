package asm

import (
	"encoding/binary"

	"github.com/J-Obog/tcvm/pkg/slf"
)

func LinkPrograms(progs []*slf.Program) *slf.Program {
	if len(progs) == 0 {
		panic("Not enough inputs to link")
	}

	base := progs[0]

	for i := 1; i < len(progs); i++ {
		link(base, progs[i])
		applyRelocs(base)
	}

	return base
}

func checkFlag(flags uint8, flag uint8) bool {
	return ((flags >> (flag - 1)) & 0x1) == 1
}

func link(prog1 *slf.Program, prog2 *slf.Program) {
	//merging symbol tables 
	for l, s2 := range prog2.SymTab {
		s1 := prog1.SymTab[l]

		if s1 == nil {
			sym := &slf.Symbol{}
			prog1.StrTab = append(prog1.StrTab, l)
			sym.StrTabIndex = uint32(len(prog1.StrTab) - 1)

			if checkFlag(s2.Flags, slf.S_ISDATA) {
				sym.Offset = s2.Offset + prog1.DataSegSize
			} else {
				sym.Offset = s2.Offset + prog1.CodeSegSize
			}
			prog1.SymTab[l] = sym
		} else {
			s1ExternFlg := checkFlag(s1.Flags, slf.S_ISEXTERN)
			s2ExternFlg := checkFlag(s2.Flags, slf.S_ISEXTERN)
			
			if !s1ExternFlg && !s2ExternFlg {
				panic("Redefinition of symbol")
			} 
			if s1ExternFlg && !s2ExternFlg {
				if checkFlag(s2.Flags, slf.S_ISDATA) {
					s1.Offset = s2.Offset + prog1.DataSegSize
				} else {
					s1.Offset = s2.Offset + prog1.CodeSegSize
				}

				s1.Flags |= (0 << slf.S_ISEXTERN)
			}
		}

		//merge reloc targets
		for _, t := range prog2.RelTab {
			lbl := prog2.StrTab[t.StrTabIndex]
			t.StrTabIndex = prog1.SymTab[lbl].StrTabIndex
			t.Offset += prog1.CodeSegSize
		} 

		prog1.RelTab = append(prog1.RelTab, prog2.RelTab...)

		
		//merge segments
		prog1.CodeSegSize += prog2.CodeSegSize
		prog1.DataSegSize += prog2.DataSegSize
		prog1.CodeSeg = append(prog1.CodeSeg, prog2.CodeSeg...)
		prog1.DataSeg = append(prog1.DataSeg, prog2.DataSeg...)
	}
}


func applyRelocs(prog *slf.Program) {
	//transform offsets into absolute addresses
	for l,s := range prog.SymTab {
		if checkFlag(s.Flags, slf.S_ISEXTERN) {
			panic("Unresolved symbol")
		}

		s.Offset += prog.EntryPoint

		if checkFlag(s.Flags, slf.S_ISDATA) {
			s.Offset += prog.CodeSegSize
		}
		
		if l == "__start__" {
			prog.StartAddress = s.Offset
		}
	}

	//actually applying the relocations
	for _, t := range prog.RelTab {
		codeOff := t.Offset
		lbl := prog.StrTab[t.StrTabIndex]
		addr := prog.SymTab[lbl].Offset
		binVal := make([]byte, 4)
		binary.BigEndian.PutUint32(binVal, addr)
		copy(prog.CodeSeg[codeOff : codeOff + 4], binVal)
	}
}

