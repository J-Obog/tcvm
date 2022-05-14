package com

//registers
const (
	R0 uint8 = 0x0
	R1 uint8 = 0x1
	R2 uint8 = 0x2
	R3 uint8 = 0x3
	R4 uint8 = 0x4
	R5 uint8 = 0x5
	R6 uint8 = 0x6
	R7 uint8 = 0x7
	R8 uint8 = 0x8
	R9 uint8 = 0x9
	R10 uint8 = 0xA
	R11 uint8 = 0xB
	R12 uint8 = 0xC
	R13 uint8 = 0xD
	R14 uint8 = 0xE
	SP uint8 = 0xF 
)

//register -> string
var REGISTER_MAPPER = map[uint8]string{
	R0: "r0",
	R1: "r1",
	R2: "r2",
	R3: "r3",
	R4: "r4",
	R5: "r5",
	R6: "r6",
	R7: "r7",
	R8: "r8",
	R9: "r9",
	R10: "r10",
	R11: "r11",
	R12: "r12",
	R13: "r13",
	R14: "r14",
	SP: "sp",
}

//string -> register
var REGISTER_TBL = map[string]uint8{
	"r0": R0,
	"r1": R1,
	"r2": R2,
	"r3": R3,
	"r4": R4,
	"r5": R5,
	"r6": R6,
	"r7": R7,
	"r8": R8,
	"r9": R9,
	"r10": R10,
	"r11": R11,
	"r12": R12,
	"r13": R13,
	"r14": R14,
	"sp": SP,
}