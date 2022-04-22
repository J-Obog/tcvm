package tcvm

import (
	"errors"
	"math"
	"os"
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