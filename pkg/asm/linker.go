package asm

import (
	"bytes"
	"os"

	"github.com/J-Obog/tcvm/pkg/slf"
)

type Linker struct {
	programs []*slf.Program
}

func NewLinker(inputs []string, out string) *Linker {
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

