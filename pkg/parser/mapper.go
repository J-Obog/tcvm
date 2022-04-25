package parser

var registers = map[string]uint8{
	"r0":  0,
	"r1":  1,
	"r2":  2,
	"r3":  3,
	"r4":  4,
	"r5":  5,
	"r6":  6,
	"r7":  7,
	"sp":  9,
	"rar": 10,
}

var dataTypes = map[string]uint8{
	"b":  1,
	"w":  2,
	"dw": 4,
}

type OpcodeMap map[string]struct {
	Number uint8
	Arity  uint8
}

var opcodes = OpcodeMap{
	"nop":    {0, 0},
	"mov8":   {1, 2},
	"mov16":  {2, 2},
	"mov32":  {3, 2},
	"add":    {4, 2},
	"sub":    {5, 2},
	"mul":    {6, 2},
	"div":    {7, 2},
	"and":    {8, 2},
	"or":     {9, 2},
	"not":    {10, 1},
	"xor":    {11, 2},
	"jmp":    {12, 1},
	"jz":     {13, 1},
	"jnz":    {14, 1},
	"jn":     {15, 1},
	"jnn":    {16, 1},
	"jp":     {17, 1},
	"jnp":    {18, 1},
	"push8":  {19, 1},
	"push16": {20, 1},
	"push32": {21, 1},
	"pop8":   {22, 0},
	"pop16":  {23, 0},
	"pop32":  {24, 0},
	"call":   {25, 1},
	"ret":    {26, 0},
	"shl":    {27, 2},
	"shr":    {28, 2},
	"sys":    {29, 0},
}