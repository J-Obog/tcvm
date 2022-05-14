package vm

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
func (c *Cpu) transferOp(dir uint8, imm uint8, size uint8, ind uint8) {
	reg1 := c.Memory[c.PC]
	var op2 uint32
	c.PC++

	if imm == 0 {
		reg2 := c.Memory[c.PC]
		op2 = c.Registers[reg2]
		c.PC++
	} else {
		op2 = stou32(SZ_DWORD, c.Memory[c.PC:])
		c.PC += 4
	}

	if ind == 0 {
		c.Registers[reg1] = stou32(SZ_DWORD, u32tos(size, op2))
	} else {
		if (op2 < c.DSP) || (op2 >= MAX_MEM_SIZE) {
			panic("Segmentation fault")
		}
		if dir == 0 {
			copy(c.Memory[op2:], u32tos(size, c.Registers[reg1]))
		} else {
			c.Registers[reg1] = stou32(size, c.Memory[op2:])
		}
	}
}

//arithmetic/logic operation
func (c *Cpu) aluOp(fn uint8, imm uint8) {
	dreg := c.Memory[c.PC]    //destination register
	dval := c.Registers[dreg] //value in destination
	var sval uint32           //value in source
	var tmp uint32
	c.PC++

	if imm == 0 {
		sval = stou32(SZ_DWORD, c.Memory[c.PC:])
		c.PC += 4
	} else {
		sreg := c.Memory[c.PC]
		sval = c.Registers[sreg]
		c.PC++
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

	c.updateFlags(tmp)

	if fn != F_CMP {
		c.Registers[dreg] = tmp
	}
}

//jump operation
func (c *Cpu) jumpOp(cond uint8, imm uint8, ret uint8) {
	if ret == 1 {
		c.PC = c.RAR
		return
	}

	var addr uint32
	var ftest bool

	if imm == 0 {
		reg := c.Memory[c.PC]
		addr = c.Registers[reg]
		c.PC++
	} else {
		addr = stou32(SZ_DWORD, c.Memory[c.PC:])
		c.PC += 4
	}

	switch cond {
	case C_UNC, C_LINK:
		ftest = true
	case C_ZERO:
		ftest = c.getFlag(FLG_ZERO)
	case C_NZERO:
		ftest = !c.getFlag(FLG_ZERO)
	case C_POS:
		ftest = c.getFlag(FLG_POS)
	case C_NPOS:
		ftest = !c.getFlag(FLG_POS)
	case C_SGN:
		ftest = c.getFlag(FLG_NEG)
	case C_NSGN:
		ftest = !c.getFlag(FLG_NEG)
	}

	if cond == C_LINK {
		c.RAR = c.PC
	}

	if ftest {
		if (addr < c.CSP) || (addr >= c.DSP) {
			panic("Segmentation fault")
		}
		c.PC = addr
	}
}
