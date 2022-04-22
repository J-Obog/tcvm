package tcvm

import (
	"errors"
	"math"
	"os"
)

const ( // register mapping
	R0 = 0 // RX = general purpose register
	R1 = 1
	R2 = 2 
	R3 = 3 
	PC = 4 // program counter
	SP = 5 // stack pointer
	RAR = 6 // return address register 
	IR = 7 // instruction register
)

type VirtualMachine struct {
	RegisterFile [16]uint32
	Memory       [math.MaxInt32]uint8
}

func (vm *VirtualMachine) LoadFromFile(path string) (error) {
	content, err := os.ReadFile(path) 
	
	if err != nil { 
		return err
	}

	if len(content) > math.MaxInt32 {
		return errors.New("Program size too big")
	}

	for i, b := range content {
		vm.Memory[i] = b
	}

	return nil
}
