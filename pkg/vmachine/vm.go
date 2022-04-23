package vmachine

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
	r4 
	r5
	r6
	r7
	pc // program counter
	sp // stack pointer
	rar // return address register 
	ir // instruction register
	dsp // data segement pointer
	sbp // stack base pointer
	hp // heap pointer
	flg // flags [HALT | ZERO | NEG]
)

const ( // status flag mapping
	hf = iota
	zf
	nf
)

const ( // sys call mapping, sys calls use r5 for mapping and subsequent registers for operands
	halt = iota
	puts
	gets
)

type VM struct {
	reg [flg + 1]uint32
	mem [math.MaxInt32]uint8 //little endian
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

func (vm *VM) write(loc uint32, sz uint8, data uint32) {
	for i := uint8(0); i < sz; i++ {
		vm.mem[loc + uint32(i)] = uint8(data >> (8*(sz - 1)))
	} 
}

func (vm *VM) updateFlags(r uint8) {
	vm.reg[flg] |= ((vm.reg[r] >> 31) << nf) // set negative flag
	
	if(vm.reg[r] == 0) { // set zero flag
		vm.reg[flg] |= (1 << zf) 
	}
}

func (vm *VM) checkFlag(flag uint8) bool {
	return (((1 << flag) & vm.reg[flg]) >> flag) == 1
}

func (vm *VM) fetch(n uint8) {
	if n > 0 {
		for i := uint8(0); i < n; i++ {
			vm.reg[ir] <<= 8
			vm.reg[ir] |= uint32(vm.mem[vm.reg[pc]])
			vm.reg[pc]++
		}
	}
}

func (vm *VM) Run() {
	for {
		if vm.checkFlag(hf) {
			break //break if halt flag is set to 1
		}

		//fetch
		vm.fetch(1)

		//decode
		opc := 63 & (vm.reg[ir] >> 2)
		mod := uint8(3 & vm.reg[ir])
		
		//execute
		opLookup[opc](vm, mod)
	}
}