package vm

import (
	"github.com/J-Obog/tcvm/pkg/com"
)

const MAX_MEM_SIZE = (1 << 16)
const REGFILE_SIZE = (1 << 4)
const ( //status flag mapping
	FLG_ZERO uint8 = 0
	FLG_NEG  uint8 = 1
	FLG_POS  uint8 = 2
)

type Cpu struct {
	//register file
	Registers [REGFILE_SIZE]uint32
	
	//memory big endian
	Memory [MAX_MEM_SIZE]uint8

	//program counter
	PC uint32

	//status flags
	Flags uint8 

	//data segment pointer
	DSP uint32

	//stack base pointer
	SBP uint32 

	//code segment pointer
	CSP uint32

	//extra segment pointer
	ESP uint32

	//return address register
	RAR uint32 
}


func (c *Cpu) memRead(addr uint32, rsize uint8) uint32 {
	ptr := addr
	end := ptr + uint32(rsize)
	data := uint32(0)

	for ptr < end {
		data <<= 8
		data |= uint32(c.Memory[ptr])
		ptr++
	}

	return data
}

func (c *Cpu) memWrite(addr uint32, wsize uint8, data uint32) {
	ptr := addr
	end := ptr + uint32(wsize)
	sfac := wsize - 1

	for ptr < end {
		word := (data & (255 << (8 * sfac))) >> (8 * sfac)
		c.Memory[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

func (c *Cpu) regRead(reg uint8, rsize uint8) uint32 {
	return c.Registers[reg] & ((1 << (8 * rsize)) - 1)
}

func (c *Cpu) regWrite(reg uint8, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	c.Registers[reg] = d
}


func (c *Cpu) updateFlags(val uint32) {
	if(val == 0) { // set zero flag
		c.Flags |= (1 << FLG_ZERO) 
	} else {
		c.Flags |= (0 << FLG_ZERO) 
	}

	sgn := val >> 31

	if sgn == 0 { //set sign flags
		c.Flags |= (1 << FLG_POS)
		c.Flags |= (0 << FLG_NEG)
	} else {
		c.Flags |= (0 << FLG_POS)
		c.Flags |= (1 << FLG_NEG)
	}
}

func (c *Cpu) getFlag(flag uint8) bool {
	return (((1 << flag) & c.Flags) >> flag) == 1
}


func (c *Cpu) Run() {
	for {
		//fetch
		op := c.Memory[c.PC]  
		primaryOp := (op >> 5) & 0x7
		c.PC++
		
		//decode/execute
		switch primaryOp {
		case com.NO_OPERATION_OP:
			//no operation

		case com.DATA_TRANSFER_OP:
			d := (op >> 4) & 0x1
			i := (op >> 3) & 0x1
			s := (op >> 1) & 0x3
			ind := op & 0x1
			c.transferOp(d, i, s, ind)

		case com.ARITHMETIC_LOGIC_OP:
			f := (op >> 1) & 0xF
			i := op & 0x1
			c.aluOp(f, i)

		case com.JUMP_OP:
			cnd := (op >> 2) & 0x7 
			i := (op >> 1) & 0x1
			r := op & 0x1 
			c.jumpOp(cnd, i, r)

		case com.SYSCALL_OP:
			c.sysCall(c.Registers[com.R0])
		}
	}
}