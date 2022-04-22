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

const ( // opcode mapping
	nop = iota 
	mov8
	mov16
	mov32
	add
	sub
	mul
	div
	and
	or
	not
	xor
	cmp
	jmp
	jz
	jnz
	jc
	jnc
	push8
	push16
	push32
	pop8
	pop16
	pop32
	call
	ret
	shl
	shr
	sprnt
	halt
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

func (vm *VM) fetch(n uint8) {
	var i uint8; 
	for i = 0; i < n; i++ {
		vm.reg[ir] <<= 8
		vm.reg[ir] |= uint32(vm.mem[vm.reg[pc]])
		vm.reg[pc]++
	}
}