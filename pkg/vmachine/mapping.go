package vmachine

const MAX_MEM_SIZE = (1 << 16)

//opcodes
const OP_NOP uint8 = 0
const OP_MOV uint8 = 1
const OP_ADD uint8 = 2
const OP_SUB uint8 = 3
const OP_MUL uint8 = 4
const OP_DIV uint8 = 5
const OP_AND uint8 = 6
const OP_OR uint8 = 7
const OP_NOT uint8 = 8
const OP_XOR uint8 = 9
const OP_CMP uint8 = 10
const OP_JMP uint8 = 11
const OP_PUSH uint8 = 12
const OP_POP uint8 = 13
const OP_CALL uint8 = 14
const OP_RET uint8 = 15
const OP_SHL uint8 = 16
const OP_SHR uint8 = 17
const OP_SYS uint8 = 18

// status flags
const F_HALT uint8 = 0
const F_ZERO uint8 = 1
const F_NEG uint8 = 2
const F_POS uint8 = 3

//system calls
const SYS_HALT uint8 = 0
const SYS_PUTS uint8 = 1
const SYS_GETS uint8 = 2

//registers
const R_R0 uint8 = 0
const R_R1 uint8 = 1
const R_R2 uint8 = 2
const R_R3 uint8 = 3
const R_R4 uint8 = 4
const R_R5 uint8 = 5
const R_R6 uint8 = 6
const R_R7 uint8 = 7
const R_SP uint8 = 8

//address modes
const M_REG uint8 = 0
const M_EREG uint8 = 1
const M_EMEM uint8 = 2
const M_IMMED uint8 = 3
const M_MEM uint8 = 4

//data sizes
const BYTE uint8 = 1
const WORD uint8 = 2
const DWORD uint8 = 4