package vm

import "github.com/J-Obog/tcvm/pkg/com"

const ( //function mapping for alu operations
	F_ADD uint8 = 0
	F_SUB uint8 = 1
	F_MUL uint8 = 2
	F_DIV uint8 = 3
	F_AND uint8 = 4
	F_OR  uint8 = 5
	F_XOR uint8 = 6
	F_NOT uint8 = 7
	F_CMP uint8 = 8
	F_SHL uint8 = 9
	F_SHR uint8 = 10
)

const ( //jump condition mapping
	C_UNC   uint8 = 0
	C_ZERO  uint8 = 1
	C_NZERO uint8 = 2
	C_POS   uint8 = 3
	C_NPOS  uint8 = 4
	C_SGN   uint8 = 5
	C_NSGN  uint8 = 6
	C_LINK  uint8 = 7
)

//data transfer operation
func (m *VirtualMachine) transferOp(dir uint8, imm uint8, size uint8, ind uint8) {
	reg1 := m.Memory[m.PC]
	var reg2 uint8
	var op2 uint32
	szMap := [3]uint8{1, 2, 4}
	sz := szMap[size]
	m.PC++

	if imm == 0 {
		reg2 = m.Memory[m.PC]
		op2 = m.Registers[reg2]
		m.PC++
	} else {
		op2 = m.memRead(m.PC, 0x4)
		m.PC += 4
	}

	if ind == 0 {
		m.regWrite(reg1, sz, op2)
	} else {
		if op2 < m.DSP {
			panic("Segmentation fault")
		}

		if dir == 0 {
			if (imm == 0) && (reg2 == com.SP) {
				if op2 < m.SBP {
					panic("Stack underflow")
				}
				eAddr := ((op2 + uint32(sz)) - 1)
				if (eAddr >= MAX_MEM_SIZE) || (eAddr >= m.ESP) {
					panic("Stack overflow")
				}
			}

			m.memWrite(op2, sz, m.Registers[reg1])
		} else {
			m.Registers[reg1] = m.memRead(op2, sz)
		}
	}
}

//arithmetic/logic operation
func (m *VirtualMachine) aluOp(fn uint8, imm uint8) {
	dreg := m.Memory[m.PC] //destination register
	dval := m.Registers[dreg] //value in destination
	var sval uint32       //value in source
	var tmp uint32
	m.PC++

	if imm == 0 {
		sval = m.memRead(m.PC, 0x4)
		m.PC += 4
	} else {
		sreg := m.Memory[m.PC]
		sval = m.Registers[sreg]
		m.PC++
	}

	switch fn {
	case F_ADD:
		tmp = dval + sval
	case F_SUB, F_CMP:
		tmp = dval + ((^sval) + 1)
	case F_MUL:
		tmp = dval * sval
	case F_DIV:
		tmp = dval / sval
	case F_AND:
		tmp = dval & sval
	case F_OR:
		tmp = dval | sval
	case F_XOR:
		tmp = dval ^ sval
	case F_NOT:
		tmp = ^sval
	case F_SHL:
		tmp = dval << sval
	case F_SHR:
		tmp = dval >> sval
	}

	m.updateFlags(tmp)

	if fn != F_CMP {
		m.Registers[dreg] = tmp
	}
}

//jump operation
func (m *VirtualMachine) jumpOp(cond uint8, imm uint8, ret uint8) {
	if ret == 1 {
		m.PC = m.RAR
		return
	}

	var addr uint32
	var ftest bool

	if imm == 0 {
		reg := m.Memory[m.PC]
		addr = m.Registers[reg]
		m.PC++
	} else {
		addr = m.memRead(m.PC, 0x4)
		m.PC += 4
	}

	switch cond {
	case C_UNC, C_LINK:
		ftest = true
	case C_ZERO:
		ftest = m.getFlag(FLG_ZERO)
	case C_NZERO:
		ftest = !m.getFlag(FLG_ZERO)
	case C_POS:
		ftest = m.getFlag(FLG_POS)
	case C_NPOS:
		ftest = !m.getFlag(FLG_POS)
	case C_SGN:
		ftest = m.getFlag(FLG_NEG)
	case C_NSGN:
		ftest = !m.getFlag(FLG_NEG)
	}

	if cond == C_LINK {
		m.RAR = m.PC
	}

	if ftest {
		if (addr < m.CSP) || (addr >= m.DSP) {
			panic("Segmentation fault")
		}
		m.PC = addr
	}
}
