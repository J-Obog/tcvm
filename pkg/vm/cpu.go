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

type VirtualMachine struct {
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


func (m *VirtualMachine) memRead(addr uint32, rsize uint8) uint32 {
	ptr := addr
	end := ptr + uint32(rsize)
	data := uint32(0)

	for ptr < end {
		data <<= 8
		data |= uint32(m.Memory[ptr])
		ptr++
	}

	return data
}

func (m *VirtualMachine) memWrite(addr uint32, wsize uint8, data uint32) {
	ptr := addr
	end := ptr + uint32(wsize)
	sfac := wsize - 1

	for ptr < end {
		word := (data & (255 << (8 * sfac))) >> (8 * sfac)
		m.Memory[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

func (m *VirtualMachine) regRead(reg uint8, rsize uint8) uint32 {
	return m.Registers[reg] & ((1 << (8 * rsize)) - 1)
}

func (m *VirtualMachine) regWrite(reg uint8, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	m.Registers[reg] = d
}


func (m *VirtualMachine) updateFlags(val uint32) {
	if(val == 0) { // set zero flag
		m.Flags |= (1 << FLG_ZERO) 
	} else {
		m.Flags |= (0 << FLG_ZERO) 
	}

	sgn := val >> 31

	if sgn == 0 { //set sign flags
		m.Flags |= (1 << FLG_POS)
		m.Flags |= (0 << FLG_NEG)
	} else {
		m.Flags |= (0 << FLG_POS)
		m.Flags |= (1 << FLG_NEG)
	}
}

func (m *VirtualMachine) getFlag(flag uint8) bool {
	return (((1 << flag) & m.Flags) >> flag) == 1
}


func (m *VirtualMachine) Run() {
	for {
		//fetch
		op := m.Memory[m.PC]  
		primaryOp := (op >> 5) & 0x7
		m.PC++
		
		//decode/execute
		switch primaryOp {
		case com.NO_OPERATION_OP:
			//no operation

		case com.DATA_TRANSFER_OP:
			d := (op >> 4) & 0x1
			i := (op >> 3) & 0x1
			s := (op >> 1) & 0x3
			ind := op & 0x1
			m.transferOp(d, i, s, ind)

		case com.ARITHMETIC_LOGIC_OP:
			f := (op >> 1) & 0xF
			i := op & 0x1
			m.aluOp(f, i)

		case com.JUMP_OP:
			c := (op >> 2) & 0x7 
			i := (op >> 1) & 0x1
			r := op & 0x1 
			m.jumpOp(c, i, r)

		case com.SYSCALL_OP:
			m.sysCall(m.Registers[com.R0])
		}
	}
}