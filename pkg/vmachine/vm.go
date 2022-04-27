package vmachine

import (
	"errors"
	"os"
)


type VM struct {
	//register file
	regf [8][4]byte
	
	//memory big endian
	mem [MAX_MEM_SIZE]byte 
	
	//program counter
	pc uint8 

	//status flags
	flags byte 

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

	for i, b := range content {
		vm.mem[i] = b
	}

	return nil
}

/*func (vm *VM) write(loc uint32, sz uint8, data uint32) {
	for i := uint8(0); i < sz; i++ {
		vm.mem[loc + uint32(i)] = uint8(data >> (8*(sz - 1)))
	} 
}

func (vm *VM) updateFlags(r uint8) {
	sgn := (vm.reg[r] >> 31)
	vm.reg[flg] |= (sgn << nf) // set negative flag
	vm.reg[flg] |= ((1 ^ sgn) << pf) // set positive flag

	if(vm.reg[r] == 0) { // set zero flag
		vm.reg[flg] |= (1 << zf) 
	}
}

func (vm *VM) checkFlag(flag uint8) bool {
	return (((1 << flag) & vm.reg[flg]) >> flag) == 1
}

func (vm *VM) fetch(n uint8) {
	if n > 0 {
		for i := uint8(0); i < n; i++ {
			vm.reg[ir] <<= 8
			vm.reg[ir] |= uint32(vm.mem[vm.reg[pc]])
			vm.reg[pc]++
		}
	}
}*/


func (vm *VM) Run() {
	for {
		/*if vm.checkFlag(hf) {
			break //break if halt flag is set to 1
		}*/

		//fetch instruction
		instruction := vm.mem[vm.pc]
		vm.pc++

		//fetch operands
		operands := vm.mem[vm.pc]
		vm.pc++

		//decode
		opc := 63 & (instruction >> 2)
		suff := 3 & instruction
		dest := operands >> 4 
		src := 15 & operands

		//execute
		opLookup[opc](vm, suff, dest, src)
	}
}