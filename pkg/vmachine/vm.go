package vmachine

import (
	"errors"
	"os"
)

const MAX_MEM_SIZE = (1 << 16)
const REGFILE_SIZE = (1 << 4)

type VM struct {
	//register file
	regs [REGFILE_SIZE]uint32
	
	//memory big endian
	ram [MAX_MEM_SIZE]uint8

	//program counter
	pc uint32

	//status flags
	flags uint8 

	//data segment pointer
	dsp uint32

	//stack base pointer
	sbp uint32 

	//code segment pointer
	csp uint32

	//extra segment pointer
	esp uint32

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


func (vm *VM) updateFlags(val uint32) {
	if(val == 0) { // set zero flag
		vm.flags |= (1 << FLG_ZERO) 
	} else {
		vm.flags |= (0 << FLG_ZERO) 
	}

	sgn := val >> 31

	if sgn == 0 { //set sign flags
		vm.flags |= (1 << FLG_POS)
		vm.flags |= (0 << FLG_NEG)
	} else {
		vm.flags |= (0 << FLG_POS)
		vm.flags |= (1 << FLG_NEG)
	}
}

func (vm *VM) getFlag(flag uint8) bool {
	return (((1 << flag) & vm.flags) >> flag) == 1
}


func (vm *VM) Run() {
	for {
		if vm.getFlag(FLG_HALT) {
			break
		}
		
		//fetch
		op := vm.ram[vm.pc]  
		primaryOp := (op >> 5) & 0x7
		vm.pc++
		
		//decode/execute
		switch primaryOp {
		case OP_NOP:
			//no operation

		case OP_DT:
			d := (op >> 4) & 0x1
			i := (op >> 3) & 0x1
			s := (op >> 1) & 0x3
			ind := op & 0x1
			vm.transferOp(d, i, s, ind)

		case OP_ALU:
			f := (op >> 1) & 0xF
			i := op & 0x1
			vm.aluOp(f, i)

		case OP_JMP:
			c := (op >> 2) & 0x7 
			i := (op >> 1) & 0x1
			r := op & 0x1 
			vm.jumpOp(c, i, r)

		case OP_SYS:
			vm.sysCall(vm.regs[R2])
		}
	}
}