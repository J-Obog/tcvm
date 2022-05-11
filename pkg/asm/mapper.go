package asm

var registers = map[string]uint8{
	"r0": 0x0,
	"r1": 0x1,
	"r2": 0x2,
	"r3": 0x3,
	"r4": 0x4,
	"r5": 0x5,
	"r6": 0x6,
	"r7": 0x7,
	"sp": 0x8,
}

const ( //op type mapping
	DTransfer uint8 = 0
	Alu       uint8 = 1
	Jump      uint8 = 2
	Nop       uint8 = 3
	SysCall   uint8 = 4
)

const ( //operand source type mapping
	Register  uint8 = 0
	Memory    uint8 = 1
	Immediate uint8 = 2
)

const ( //data def specifier
	Byte  uint8 = 0
	Word  uint8 = 1
	DWord uint8 = 2
	Space uint8 = 3
)

var dataSpecMap = map[string]uint8{
	"byte":  Byte,
	"word":  Word,
	"dword": DWord,
	"space": Space,
}

var opcodeMap = map[string]uint8{
	"mov":   0x4,
	"movb":  0x0,
	"movw":  0x2,
	"movi":  0xC,
	"movbi": 0x8,
	"movwi": 0xA,
	"ld":    0x15,
	"ldb":   0x11,
	"ldw":   0x13,
	"ldi":   0x1D,
	"ldbi":  0x19,
	"ldwi":  0x1B,
	"st":    0x5,
	"stb":   0x1,
	"stw":   0x3,
	"sti":   0xD,
	"stbi":  0x9,
	"stwi":  0xB,
	"add":   0x20,
	"sub":   0x22,
	"mul":   0x24,
	"div":   0x26,
	"and":   0x28,
	"or":    0x2A,
	"xor":   0x2C,
	"not":   0x2E,
	"cmp":   0x30,
	"shl":   0x32,
	"shr":   0x34,
	"addi":  0x21,
	"subi":  0x23,
	"muli":  0x25,
	"divi":  0x27,
	"andi":  0x29,
	"ori":   0x2B,
	"xori":  0x2D,
	"noti":  0x2F,
	"cmpi":  0x31,
	"shli":  0x33,
	"shri":  0x35,
	"jmp":   0x40,
	"jz":    0x44,
	"jnz":   0x48,
	"jp":    0x4C,
	"jnp":   0x50,
	"js":    0x54,
	"jns":   0x58,
	"jal":   0x5C,
	"jr":    0x41,
	"jmpi":  0x42,
	"jzi":   0x46,
	"jnzi":  0x4A,
	"jpi":   0x4E,
	"jnpi":  0x52,
	"jsi":   0x56,
	"jnsi":  0x5A,
	"jali":  0x5E,
	"nop":   0x60,
	"sys":   0x80,
}