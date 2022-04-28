package vmachine

import (
	"errors"
	"os"
)


type VM struct {
	//register file
	regs [8]uint32
	
	//memory big endian
	mem [MAX_MEM_SIZE]uint8 
	
	//program counter
	pc uint32

	//status flags
	flags uint8 

	//stack base pointer
	sbp uint32 

	//return address register
	rar uint32 
}

func (vm *VM) LoadFromFile(path string) (error) {
	content, err := os.ReadFile(path) 

	if err != nil { 
		return err
	}

	if len(content) > MAX_MEM_SIZE {
		return errors.New("Program size too big")
	}

	for i, b := range content {
		vm.mem[i] = b
	}

	return nil
}

/*
func (vm *VM) updateFlags(r uint8) {
	sgn := (vm.reg[r] >> 31)
	vm.reg[flg] |= (sgn << nf) // set negative flag
	vm.reg[flg] |= ((1 ^ sgn) << pf) // set positive flag

	if(vm.reg[r] == 0) { // set zero flag
		vm.reg[flg] |= (1 << zf) 
	}
}

}*/

func (vm *VM) mem_read(addr uint32, rsize uint8) uint32 {
	ptr := addr
	end := ptr + uint32(rsize)
	data := uint32(0)
	
	for ptr < end {
		data <<= 8
		data |= uint32(vm.mem[ptr])
		ptr++
	}

	return data
}

func (vm *VM) mem_write(addr uint32, wsize uint8, data uint32) {
	ptr := addr
	end := ptr + uint32(wsize)
	sfac := wsize - 1

	for ptr < end {
		word := (data & (255 << (8 * sfac))) >> (8 * sfac)
		vm.mem[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

func (vm *VM) reg_read(reg uint32, rsize uint8) uint32 {
	return vm.regs[reg] & ((1 << (8 * rsize)) - 1)
}

func (vm *VM) reg_write(reg uint32, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	vm.regs[reg] |= d
}


func (vm *VM) getFlag(flag uint8) bool {
	return (((1 << flag) & vm.flags) >> flag) == 1
}

func (vm *VM) Run() {
	for {
		if vm.getFlag(F_HALT) {
			break
		}

		//fetch instruction
		instruction := vm.mem[vm.pc]
		vm.pc++

		//fetch operand header
		operands := vm.mem[vm.pc]
		vm.pc++

		//decode
		opc := 63 & (instruction >> 2)
		suff := 3 & instruction
		dest := operands >> 4 
		src := 15 & operands

		//execute
		opLookup[opc](vm, suff, dest, src)
	}
}