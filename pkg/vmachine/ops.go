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

//data transfer operation
func (vm *VM) transferOp(dir uint8, imm uint8, size uint8, ind uint8) {

}

//arithmetic/logic operation
func (vm *VM) aluOp(fn uint8, imm uint8) {
	regs := vm.ram[vm.pc]
	vm.pc++

	r := regs & 0x7
	v1 := vm.regs[r]
	var v2 uint32

	if imm == 0 {
		v2 = vm.regs[(regs>>3)&0x7]
	} else {
		v2 = vm.memRead(vm.pc, 0x4)
		vm.pc += 4
	}

	switch fn {
	case F_ADD:
		v2 = v1 + v2
	case F_SUB, F_CMP:
		v2 = v1 + ((^v2) + 1)
	case F_MUL:
		v2 = v1 * v2
	case F_DIV:
		v2 = v1 / v2
	case F_AND:
		v2 = v1 & v2
	case F_OR:
		v2 = v1 | v2
	case F_XOR:
		v2 = v1 ^ v2
	case F_NOT:
		v2 = ^v2
	case F_SHL:
		v2 = v1 << v2
	case F_SHR:
		v2 = v1 >> v2
	}

	vm.updateFlags(v2)

	if fn != F_CMP {
		vm.regs[r] = v2
	}
}

//jump operation
func (vm *VM) jumpOp(cond uint8, imm uint8, ret uint8) {

}

//system call
func (vm *VM) sysCall(num uint32) {

}