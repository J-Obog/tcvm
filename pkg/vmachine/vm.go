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
	flg // flags [HALT | ZERO | CARRY]
)

const ( // status flag mapping
	hf = iota
	zf
	cf
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

func subBits(x uint32, k uint8, p uint8) (uint32) {
	return ((1 << k) - 1) & (x >> (p - 1))
}

func kthBit (x uint32, k uint8) (uint8) {
	return uint8((x & (1 << (k - 1))) >> (k - 1))
}


func (vm *VM) setBlock (p uint32, d uint32, s uint32) {
	for p < p+s {
		vm.mem[p] = uint8(d)
		p++;
		d >>= 8
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
		hs := kthBit(vm.reg[flg], hf) //get halt status flag
		if hs == 1 {
			break //break if halt flag is set to 1
		}

		//fetch
		vm.fetch(1)

		//decode
		opc := uint8(subBits(vm.reg[ir], 6, 3))
		opr := uint8(subBits(vm.reg[ir], 2, 1))
		
		//execute
		opLookup[opc](vm, opr)
	}
}