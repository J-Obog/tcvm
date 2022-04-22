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
	pc // program counter
	sp // stack pointer
	rar // return address register 
	ir // instruction register
	flg // flags [HALT | ZERO | NEG]
)

const ( // status flag mapping
	hf = iota
	zf
	nf
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

func (vm *VM) setFlags(r uint8) {
	vm.reg[flg] |= ((vm.reg[r] >> 31) << nf) // set negative flag
	
	if(vm.reg[r] == 0) { // set zero flag
		vm.reg[flg] |= (1 << zf) 
	}
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
		hs := ((1 << hf) & vm.reg[flg]) >> hf//get halt status flag
		if hs == 1 {
			break //break if halt flag is set to 1
		}

		//fetch
		vm.fetch(1)

		//decode
		opc := 63 & (vm.reg[ir] >> 2)
		opr := uint8(3 & vm.reg[ir])
		
		//execute
		opLookup[opc](vm, opr)
	}
}