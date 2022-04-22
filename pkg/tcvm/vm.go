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

func (vm *VM) subBits(x uint32, k uint8, p uint8) (uint32) {
	return ((1 << k) - 1) & (x >> (p - 1))
}

func (vm *VM) fetch(n uint8) {
	var i uint8; 
	for i = 0; i < n; i++ {
		vm.reg[ir] <<= 8
		vm.reg[ir] |= uint32(vm.mem[vm.reg[pc]])
		vm.reg[pc]++
	}
}

func (vm *VM) Run() {
	for {
		vm.fetch(1)
		opc := uint8(vm.subBits(vm.reg[ir], 6, 3))
		opr := uint8(vm.subBits(vm.reg[ir], 2, 1))
		opLookup[opc](vm, opr)
	}
}