package vmachine

import (
	"errors"
	"os"
)


type VM struct {
	//register file
	registers RegisterFile
	
	//memory big endian
	ram Memory

	//program counter
	pc uint32

	//status flags
	flags uint8 

	//stack base pointer
	sbp uint32 

	//return address register
	rar uint32 
}

func (vm *VM) LoadFromFile(path string) (error) {
	content, err := os.ReadFile(path) 

	if err != nil { 
		return err
	}

	if len(content) > MAX_MEM_SIZE {
		return errors.New("Program size too big")
	}

	/*for i, b := range content {
		vm.mem[i] = b
	}*/

	return nil
}

/*
func (vm *VM) updateFlags(r uint8) {
	sgn := (vm.reg[r] >> 31)
	vm.reg[flg] |= (sgn << nf) // set negative flag
	vm.reg[flg] |= ((1 ^ sgn) << pf) // set positive flag

	if(vm.reg[r] == 0) { // set zero flag
		vm.reg[flg] |= (1 << zf) 
	}
}

}*/





func (vm *VM) getFlag(flag uint8) bool {
	return (((1 << flag) & vm.flags) >> flag) == 1
}

func (vm *VM) Run() {
	for {
		if vm.getFlag(F_HALT) {
			break
		}

		/*//fetch instruction
		instruction := vm.ram.Read(vm.pc, BYTE)
		vm.pc++

		//fetch operand header
		operands := vm.mem[vm.pc]
		vm.pc++

		//decode
		opc := 63 & (instruction >> 2)
		suff := 3 & instruction
		dest := operands >> 4 
		src := 15 & operands

		//execute
		opLookup[opc](vm, suff, dest, src)*/
	}
}