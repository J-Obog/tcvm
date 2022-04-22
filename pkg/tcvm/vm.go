package tcvm

import (
	"errors"
	"math"
	"os"
)

const ( // register mapping
	r0 = 0 // RX = general purpose register
	r1 = 1
	r2 = 2 
	r3 = 3 
	pc = 4 // program counter
	sp = 5 // stack pointer
	rar = 6 // return address register 
	ir = 7 // instruction register
	flg = 8 // flags [HALT | ZERO | CARRY]
)

type VM struct {
	reg [16]uint32
	mem [math.MaxInt32]uint8
}

func (vm *VM) LoadFromFile(path string) (error) {
	content, err := os.ReadFile(path) 
	
	if err != nil { 
		return err
	}

	if len(content) > math.MaxInt32 {
		return errors.New("Program size too big")
	}

	for i, b := range content {
		vm.mem[i] = b
	}

	return nil
}

func (vm *VM) fetch() {
	vm.reg[ir] = vm.reg[pc]
	vm.reg[pc]++
}