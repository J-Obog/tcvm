package vmachine

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

const ( //system call mapping
	S_HALT uint32 = 0
	S_PUTS uint32 = 1
	S_GETS uint32 = 2
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
	reg1 := m.ram[m.pc]
	var reg2 uint8
	var op2 uint32
	szMap := [3]uint8{1, 2, 4}
	sz := szMap[size]
	m.pc++

	if imm == 0 {
		reg2 = m.ram[m.pc]
		op2 = m.regs[reg2]
		m.pc++
	} else {
		op2 = m.memRead(m.pc, 0x4)
		m.pc += 4
	}

	if ind == 0 {
		m.regWrite(reg1, sz, op2)
	} else {
		if op2 < m.dsp {
			panic("Segmentation fault")
		}

		if dir == 0 {
			if (imm == 0) && (reg2 == com.SP) {
				if op2 < m.sbp {
					panic("Stack underflow")
				}
				eAddr := ((op2 + uint32(sz)) - 1)
				if (eAddr >= MAX_MEM_SIZE) || (eAddr >= m.esp) {
					panic("Stack overflow")
				}
			}

			m.memWrite(op2, sz, m.regs[reg1])
		} else {
			m.regs[reg1] = m.memRead(op2, sz)
		}
	}
}

//arithmetic/logic operation
func (m *VirtualMachine) aluOp(fn uint8, imm uint8) {
	dreg := m.ram[m.pc] //destination register
	dval := m.regs[dreg] //value in destination
	var sval uint32       //value in source
	var tmp uint32
	m.pc++

	if imm == 0 {
		sval = m.memRead(m.pc, 0x4)
		m.pc += 4
	} else {
		sreg := m.ram[m.pc]
		sval = m.regs[sreg]
		m.pc++
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
		m.regs[dreg] = tmp
	}
}

//jump operation
func (m *VirtualMachine) jumpOp(cond uint8, imm uint8, ret uint8) {
	if ret == 1 {
		m.pc = m.rar
		return
	}

	var addr uint32
	var ftest bool

	if imm == 0 {
		reg := m.ram[m.pc]
		addr = m.regs[reg]
		m.pc++
	} else {
		addr = m.memRead(m.pc, 0x4)
		m.pc += 4
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
		m.rar = m.pc
	}

	if ftest {
		if (addr < m.csp) || (addr >= m.dsp) {
			panic("Segmentation fault")
		}
		m.pc = addr
	}
}

//system call
func (m *VirtualMachine) sysCall(num uint32) {

}