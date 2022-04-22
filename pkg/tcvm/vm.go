package tcvm

import (
	"errors"
	"math"
	"os"
)

const ( // register mapping
	r0 = iota // RX = general purpose register
	r1 
	r2 
	r3  
	pc // program counter
	sp // stack pointer
	rar // return address register 
	ir // instruction register
	flg // flags [HALT | ZERO | CARRY]
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