package vmachine

const MAX_MEM_SIZE = (1 << 16)

//opcodes
const (
	OP_NOP  uint8 = 0
	OP_MOV        = 1
	OP_ADD        = 2
	OP_SUB        = 3
	OP_MUL        = 4
	OP_DIV        = 5
	OP_AND        = 6
	OP_OR         = 7
	OP_NOT        = 8
	OP_XOR        = 9
	OP_CMP        = 10
	OP_JMP        = 11
	OP_PUSH       = 12
	OP_POP        = 13
	OP_CALL       = 14
	OP_RET        = 15
	OP_SHL        = 16
	OP_SHR        = 17
	OP_SYS        = 18
)

// status flags

const (
	F_HALT uint8 = 0
	F_ZERO       = 1
	F_NEG        = 2
	F_POS        = 3
)

//system calls
const (
	SYS_HALT uint8 = 0
	SYS_PUTS       = 1
	SYS_GETS       = 2
)

//registers
const (
	R_R0 uint8 = 0
	R_R1       = 1
	R_R2       = 2
	R_R3       = 3
	R_R4       = 4
	R_R5       = 5
	R_R6       = 6
	R_R7       = 7
	R_SP       = 8
)

//address modes
const (
	M_REG   uint8 = 0
	M_EREG        = 1
	M_EMEM        = 2
	M_IMMED       = 3
	M_MEM         = 4
)

//data sizes
const (
	BYTE  uint8 = 1
	WORD        = 2
	DWORD       = 4
)