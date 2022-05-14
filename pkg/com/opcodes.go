package com

//primary opcodes
const (
	DATA_TRANSFER_OP  uint8 = 0
	ARITHMETIC_LOGIC_OP uint8 = 1
	JUMP_OP uint8 = 2
	NO_OPERATION_OP uint8 = 3
	SYSCALL_OP uint8 = 4
)

//opcodes
const ( 
	MOV   uint8 = 0x4
	MOVB  uint8 = 0x0
	MOVW  uint8 = 0x2
	MOVI  uint8 = 0xC
	MOVBI uint8 = 0x8
	MOVWI uint8 = 0xA
	LD    uint8 = 0x15
	LDB   uint8 = 0x11
	LDW   uint8 = 0x13
	LDI   uint8 = 0x1D
	LDBI  uint8 = 0x19
	LDWI  uint8 = 0x1B
	ST    uint8 = 0x5
	STB   uint8 = 0x1
	STW   uint8 = 0x3
	STI   uint8 = 0xD
	STBI  uint8 = 0x9
	STWI  uint8 = 0xB
	ADD   uint8 = 0x20
	SUB   uint8 = 0x22
	MUL   uint8 = 0x24
	DIV   uint8 = 0x26
	AND   uint8 = 0x28
	OR    uint8 = 0x2A
	XOR   uint8 = 0x2C
	NOT   uint8 = 0x2E
	CMP   uint8 = 0x30
	SHL   uint8 = 0x32
	SHR   uint8 = 0x34
	ADDI  uint8 = 0x21
	SUBI  uint8 = 0x23
	MULI  uint8 = 0x25
	DIVI  uint8 = 0x27
	ANDI  uint8 = 0x29
	ORI   uint8 = 0x2B
	XORI  uint8 = 0x2D
	NOTI  uint8 = 0x2F
	CMPI  uint8 = 0x31
	SHLI  uint8 = 0x33
	SHRI  uint8 = 0x35
	JMP   uint8 = 0x40
	JZ    uint8 = 0x44
	JNZ   uint8 = 0x48
	JP    uint8 = 0x4C
	JNP   uint8 = 0x50
	JS    uint8 = 0x54
	JNS   uint8 = 0x58
	JAL   uint8 = 0x5C
	JR    uint8 = 0x41
	JMPI  uint8 = 0x42
	JZI   uint8 = 0x46
	JNZI  uint8 = 0x4A
	JPI   uint8 = 0x4E
	JNPI  uint8 = 0x52
	JSI   uint8 = 0x56
	JNSI  uint8 = 0x5A
	JALI  uint8 = 0x5E
	NOP   uint8 = 0x60
	SYS   uint8 = 0x80
)

//opcode -> string
var INTSTRUCTION_MAPPER = map[uint8]string{
	MOV:   "mov",
	MOVB:  "movb",
	MOVW:  "movw",
	MOVI:  "movi",
	MOVBI: "movbi",
	MOVWI: "movwi",
	LD:    "ld",
	LDB:   "ldb",
	LDW:   "ldw",
	LDI:   "ldi",
	LDBI:  "ldbi",
	LDWI:  "ldwi",
	ST:    "st",
	STB:   "stb",
	STW:   "stw",
	STI:   "sti",
	STBI:  "stbi",
	STWI:  "stwi",
	ADD:   "add",
	SUB:   "sub",
	MUL:   "mul",
	DIV:   "div",
	AND:   "and",
	OR:    "or",
	XOR:   "xor",
	NOT:   "not",
	CMP:   "cmp",
	SHL:   "shl",
	SHR:   "shr",
	ADDI:  "addi",
	SUBI:  "subi",
	MULI:  "muli",
	DIVI:  "divi",
	ANDI:  "andi",
	ORI:   "ori",
	XORI:  "xori",
	NOTI:  "noti",
	CMPI:  "cmpi",
	SHLI:  "shli",
	SHRI:  "shri",
	JMP:   "jmp",
	JZ:    "jz",
	JNZ:   "jnz",
	JP:    "jp",
	JNP:   "jnp",
	JS:    "js",
	JNS:   "jns",
	JAL:   "jal",
	JR:    "jr",
	JMPI:  "jmpi",
	JZI:   "jzi",
	JNZI:  "jnzi",
	JPI:   "jpi",
	JNPI:  "jnpi",
	JSI:   "jsi",
	JNSI:  "jnsi",
	JALI:  "jali",
	NOP:   "nop",
	SYS:   "sys",
}

//str -> opcode
var INSTRUCTION_TBL = map[string]uint8{
	"mov":   MOV,
	"movb":  MOVB,
	"movw":  MOVW,
	"movi":  MOVI,
	"movbi": MOVBI,
	"movwi": MOVWI,
	"ld":    LD,
	"ldb":   LDB,
	"ldw":   LDW,
	"ldi":   LDI,
	"ldbi":  LDBI,
	"ldwi":  LDWI,
	"st":    ST,
	"stb":   STB,
	"stw":   STW,
	"sti":   STI,
	"stbi":  STBI,
	"stwi":  STWI,
	"add":   ADD,
	"sub":   SUB,
	"mul":   MUL,
	"div":   DIV,
	"and":   AND,
	"or":    OR,
	"xor":   XOR,
	"not":   NOT,
	"cmp":   CMP,
	"shl":   SHL,
	"shr":   SHR,
	"addi":  ADDI,
	"subi":  SUBI,
	"muli":  MULI,
	"divi":  DIVI,
	"andi":  ANDI,
	"ori":   ORI,
	"xori":  XORI,
	"noti":  NOTI,
	"cmpi":  CMPI,
	"shli":  SHLI,
	"shri":  SHRI,
	"jmp":   JMP,
	"jz":    JZ,
	"jnz":   JNZ,
	"jp":    JP,
	"jnp":   JNP,
	"js":    JS,
	"jns":   JNS,
	"jal":   JAL,
	"jr":    JR,
	"jmpi":  JMPI,
	"jzi":   JZI,
	"jnzi":  JNZI,
	"jpi":   JPI,
	"jnpi":  JNPI,
	"jsi":   JSI,
	"jnsi":  JNSI,
	"jali":  JALI,
	"nop":   NOP,
	"sys":   SYS,
}

