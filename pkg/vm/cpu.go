package vm

import (
	"errors"
	"os"

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

func (m *VirtualMachine) LoadFromFile(path string) (error) {
	content, err := os.ReadFile(path) 

	if err != nil { 
		return err
	}

	if len(content) > MAX_MEM_SIZE {
		return errors.New("Program size too big")
	}

	for i, b := range content {
		m.ram[i] = b
	}

	return nil
}


func (m *VirtualMachine) memRead(addr uint32, rsize uint8) uint32 {
	ptr := addr
	end := ptr + uint32(rsize)
	data := uint32(0)

	for ptr < end {
		data <<= 8
		data |= uint32(m.ram[ptr])
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
		m.ram[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

func (m *VirtualMachine) regRead(reg uint8, rsize uint8) uint32 {
	return m.regs[reg] & ((1 << (8 * rsize)) - 1)
}

func (m *VirtualMachine) regWrite(reg uint8, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	m.regs[reg] = d
}


func (m *VirtualMachine) updateFlags(val uint32) {
	if(val == 0) { // set zero flag
		m.flags |= (1 << FLG_ZERO) 
	} else {
		m.flags |= (0 << FLG_ZERO) 
	}

	sgn := val >> 31

	if sgn == 0 { //set sign flags
		m.flags |= (1 << FLG_POS)
		m.flags |= (0 << FLG_NEG)
	} else {
		m.flags |= (0 << FLG_POS)
		m.flags |= (1 << FLG_NEG)
	}
}

func (m *VirtualMachine) getFlag(flag uint8) bool {
	return (((1 << flag) & m.flags) >> flag) == 1
}


func (m *VirtualMachine) Run() {
	for {
		//fetch
		op := m.ram[m.pc]  
		primaryOp := (op >> 5) & 0x7
		m.pc++
		
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
			m.sysCall(m.regs[com.R0])
		}
	}
}