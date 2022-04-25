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