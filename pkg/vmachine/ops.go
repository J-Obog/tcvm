package vmachine

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
	S_HALT uint8 = 0
	S_PUTS uint8 = 1
	S_GETS uint8 = 2
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

}

//jump operation
func (vm *VM) jumpOp(cond uint8, imm uint8, ret uint8) {

}

//system call
func (vm *VM) sysCall(num uint8) {

}