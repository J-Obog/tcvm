package asm

const ( //data def specifier
	ALLOC_BYTE  uint8 = 0
	ALLOC_WORD  uint8 = 1
	ALLOC_DWORD uint8 = 2
	ALLOC_SPACE uint8 = 3
)

var REGISTER_TBL = map[string]uint8{
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

var ALLOCTYPE_TBL = map[string]uint8{
	"byte":  ALLOC_BYTE,
	"word":  ALLOC_WORD,
	"dword": ALLOC_DWORD,
	"space": ALLOC_SPACE,
}

var INSTRUCTION_TBL = map[string]uint8{
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