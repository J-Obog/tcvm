package vmachine

type opFn func(*VM, uint8)

func nop(vm *VM, opr uint8) {
	//no operation
}

func mov8(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	vm.fetch(1)
	src := uint8(vm.reg[ir])

	switch opr {
	case 0:
		vm.reg[reg] |= uint32(src)
		break

	case 1:
		vm.reg[reg] |= uint32(uint8(vm.reg[src]))
		break

	case 2:
		vm.mem[reg] = src
		break

	case 3:
		vm.mem[reg] = uint8(vm.reg[src])
		break
	}
}

func mov16(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	vm.fetch(2)
	src := uint16(vm.reg[ir])

	switch opr {
	case 0:
		vm.reg[reg] |= uint32(src)
		break

	case 1:
		vm.reg[reg] |= uint32(uint16(vm.reg[src]))
		break

	case 2:
		vm.mem[reg] = uint8(src)
		vm.mem[reg+1] = uint8(src >> 8)
		break

	case 3:
		vm.mem[reg] = uint8(vm.reg[src])
		vm.mem[reg+1] = uint8(vm.reg[src] >> 8)
		break
	}
}

func mov32(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	vm.fetch(4)
	src := vm.reg[ir]

	switch opr {
	case 0:
		vm.reg[reg] = src
		break

	case 1:
		vm.reg[reg] = vm.reg[src]
		break

	case 2:
		vm.mem[reg] = uint8(src)
		vm.mem[reg+1] = uint8(src >> 8)
		vm.mem[reg+2] = uint8(src >> 16)
		vm.mem[reg+3] = uint8(src >> 24)
		break

	case 3:
		vm.mem[reg] = uint8(vm.reg[src])
		vm.mem[reg+1] = uint8(vm.reg[src] >> 8)
		vm.mem[reg+2] = uint8(vm.reg[src] >> 16)
		vm.mem[reg+3] = uint8(vm.reg[src] >> 24)
		break
	}
}

func add(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[reg] = uint32(uint64(vm.reg[reg]) + uint64(src))
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[reg] = uint32(uint64(vm.reg[reg]) + uint64(vm.reg[src]))
	}

	vm.setFlags(reg)
}

func sub(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[reg] = uint32(uint64(vm.reg[reg]) + uint64((^src)+1))
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[reg] = uint32(uint64(vm.reg[reg]) + uint64((^vm.reg[src])+1))
	}

	vm.setFlags(reg)
}

var opLookup = [64]opFn{
	nop,
	mov8,
	mov16,
	mov32,
	add,
	sub,
}