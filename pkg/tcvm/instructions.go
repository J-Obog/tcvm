package tcvm

type opFn func(*VM, uint8)

func nop(vm *VM, opr uint8) {
	//no operation
}

var opLookup = [64]opFn{
	nop,
}