package com

//alloc types
const (
	BYTE uint8 = 0
	WORD uint8 = 1
	DWORD uint8 = 2
	SPACE uint8 = 3
)

//string -> alloctype
var ALLOCTYPE_TBL = map[string]uint8 {
	"byte": BYTE,
	"word": WORD,
	"dword": DWORD, 
	"space": SPACE,
}