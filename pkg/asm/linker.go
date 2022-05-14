package asm

import (
	"bytes"
	"os"

	"github.com/J-Obog/tcvm/pkg/slf"
)

type Linker struct {
	programs []*slf.Program
}

func NewLinker(inputs []string) *Linker {
	l := &Linker{}

	for _, input := range inputs {
		content, err := os.ReadFile(input)

		if err != nil {
			panic(err)
		}

		pgm := &slf.Program{}
		buf := bytes.NewBuffer(content)
		pgm.Decode(buf)
		l.programs = append(l.programs, pgm)
	}

	return l
} 

func (l *Linker) Link(prog1 *slf.Program, prog2 *slf.Program) {
	//link two files
}

func (l *Linker) LinkFiles(out string) {
	if len(l.programs) == 0 {
		panic("Not enough input files to link")
	}

	base := l.programs[0]

	for i := 1; i < len(l.programs); i++ {
		l.Link(base, l.programs[i])
	}

	buf := base.Encode()
	err := os.WriteFile(out, buf.Bytes(), 0777)

	if err != nil {
		panic(err)
	}
}


