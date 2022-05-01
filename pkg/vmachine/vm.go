package vmachine

import (
	"errors"
	"os"
)


type VM struct {
	//register file
	regs [8]uint32
	
	//memory big endian
	ram [MAX_MEM_SIZE]uint8

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
		vm.ram[i] = b
	}

	return nil
}


func (vm *VM) memRead(addr uint32, rsize uint8) uint32 {
	ptr := addr
	end := ptr + uint32(rsize)
	data := uint32(0)

	for ptr < end {
		data <<= 8
		data |= uint32(vm.ram[ptr])
		ptr++
	}

	return data
}

func (vm *VM) memWrite(addr uint32, wsize uint8, data uint32) {
	ptr := addr
	end := ptr + uint32(wsize)
	sfac := wsize - 1

	for ptr < end {
		word := (data & (255 << (8 * sfac))) >> (8 * sfac)
		vm.ram[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

func (vm *VM) regRead(reg uint8, rsize uint8) uint32 {
	return vm.regs[reg] & ((1 << (8 * rsize)) - 1)
}

func (vm *VM) regWrite(reg uint8, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	vm.regs[reg] = d
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

func (vm *VM) getFlag(flag uint8) bool {
	return (((1 << flag) & vm.flags) >> flag) == 1
}


func (vm *VM) Run() {
	for {
		if vm.getFlag(F_HALT) {
			break
		}

		//fetch opcode
		opcode := vm.ram[vm.pc]
		vm.pc++
	
		//execute
		opLookup[opcode](vm)
	}
}