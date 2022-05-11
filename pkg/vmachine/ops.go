package vmachine

const (
	OP_DT  uint8 = 0
	OP_ALU uint8 = 1
	OP_JMP uint8 = 2
	OP_NOP uint8 = 3
	OP_SYS uint8 = 4
)

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

const ( //status flag mapping
	FLG_HALT uint8 = 0
	FLG_ZERO uint8 = 1
	FLG_NEG  uint8 = 2
	FLG_POS  uint8 = 3
)

const ( //system call mapping
	S_HALT uint32 = 0
	S_PUTS uint32 = 1
	S_GETS uint32 = 2
)

const ( //register mapping
	R0 uint8 = 0
	R1 uint8 = 1
	R2 uint8 = 2
	R3 uint8 = 3
	R4 uint8 = 4
	R5 uint8 = 5
	R6 uint8 = 6
	R7 uint8 = 7
	SP uint8 = 8
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
func (vm *VM) transferOp(dir uint8, imm uint8, size uint8, ind uint8) {
	reg1 := vm.ram[vm.pc]
	var reg2 uint8
	var op2 uint32
	szMap := [3]uint8{1, 2, 4}
	sz := szMap[size]
	vm.pc++

	if imm == 0 {
		reg2 = vm.ram[vm.pc]
		op2 = vm.regs[reg2]
		vm.pc++
	} else {
		op2 = vm.memRead(vm.pc, 0x4)
		vm.pc += 4
	}

	if ind == 0 {
		vm.regWrite(reg1, sz, op2)
	} else {
		if op2 < vm.dsp {
			panic("Segmentation fault")
		}

		if dir == 0 {
			if (imm == 0) && (reg2 == SP) {
				if op2 < vm.sbp {
					panic("Stack underflow")
				}
				eAddr := ((op2 + uint32(sz)) - 1)
				if (eAddr >= MAX_MEM_SIZE) || (eAddr >= vm.esp) {
					panic("Stack overflow")
				}
			}

			vm.memWrite(op2, sz, vm.regs[reg1])
		} else {
			vm.regs[reg1] = vm.memRead(op2, sz)
		}
	}
}

//arithmetic/logic operation
func (vm *VM) aluOp(fn uint8, imm uint8) {
	dreg := vm.ram[vm.pc] //destination register
	dval := vm.regs[dreg] //value in destination
	var sval uint32       //value in source
	var tmp uint32
	vm.pc++

	if imm == 0 {
		sval = vm.memRead(vm.pc, 0x4)
		vm.pc += 4
	} else {
		sreg := vm.ram[vm.pc]
		sval = vm.regs[sreg]
		vm.pc++
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

	vm.updateFlags(tmp)

	if fn != F_CMP {
		vm.regs[dreg] = tmp
	}
}

//jump operation
func (vm *VM) jumpOp(cond uint8, imm uint8, ret uint8) {
	if ret == 1 {
		vm.pc = vm.rar
		return
	}

	var addr uint32
	var ftest bool

	if imm == 0 {
		reg := vm.ram[vm.pc]
		addr = vm.regs[reg]
		vm.pc++
	} else {
		addr = vm.memRead(vm.pc, 0x4)
		vm.pc += 4
	}

	switch cond {
	case C_UNC, C_LINK:
		ftest = true
	case C_ZERO:
		ftest = vm.getFlag(FLG_ZERO)
	case C_NZERO:
		ftest = !vm.getFlag(FLG_ZERO)
	case C_POS:
		ftest = vm.getFlag(FLG_POS)
	case C_NPOS:
		ftest = !vm.getFlag(FLG_POS)
	case C_SGN:
		ftest = vm.getFlag(FLG_NEG)
	case C_NSGN:
		ftest = !vm.getFlag(FLG_NEG)
	}

	if cond == C_LINK {
		vm.rar = vm.pc
	}

	if ftest {
		if (addr < vm.csp) || (addr >= vm.dsp) {
			panic("Segmentation fault")
		}
		vm.pc = addr
	}
}

//system call
func (vm *VM) sysCall(num uint32) {

}