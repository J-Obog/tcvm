package tcvm

type VirtualMachine struct {
	RegisterFile [16]uint32
	Memory       [2147483647]uint8
}