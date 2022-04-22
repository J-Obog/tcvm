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

func and(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[reg] &= src
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[reg] &= vm.reg[src]
	}

	vm.setFlags(reg)
}

func or(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[reg] |= src
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[reg] |= vm.reg[src]
	}

	vm.setFlags(reg)
}

func xor(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])

	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[reg] ^= src
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[reg] ^= vm.reg[src]
	}

	vm.setFlags(reg)
}

func not(vm *VM, opr uint8) {
	vm.fetch(1)
	reg := uint8(vm.reg[ir])
	vm.reg[reg] = ^vm.reg[reg]

	vm.setFlags(reg)
}

func jmp(vm *VM, opr uint8) {
	if opr == 0 {
		vm.fetch(4)
		src := vm.reg[ir]
		vm.reg[pc] = src
	} else {
		vm.fetch(1)
		src := uint8(vm.reg[ir])
		vm.reg[pc] = vm.reg[src]
	}
}

func jz(vm *VM, opr uint8) {
	var src uint32

	if opr == 0 {
		vm.fetch(4)
		src = vm.reg[ir]
	} else {
		vm.fetch(1)
		r := uint8(vm.reg[ir])
		src = vm.reg[r]
	}

	zfs := ((1 << zf) & vm.reg[flg]) >> zf
	if zfs == 1 {
		vm.reg[pc] = src
	}
}

func jnz(vm *VM, opr uint8) {
	var src uint32

	if opr == 0 {
		vm.fetch(4)
		src = vm.reg[ir]
	} else {
		vm.fetch(1)
		r := uint8(vm.reg[ir])
		src = vm.reg[r]
	}

	zfs := ((1 << zf) & vm.reg[flg]) >> zf
	if zfs == 0 {
		vm.reg[pc] = src
	}
}

func jn(vm *VM, opr uint8) {
	var src uint32

	if opr == 0 {
		vm.fetch(4)
		src = vm.reg[ir]
	} else {
		vm.fetch(1)
		r := uint8(vm.reg[ir])
		src = vm.reg[r]
	}

	nfs := ((1 << nf) & vm.reg[flg]) >> nf
	if nfs == 1 {
		vm.reg[pc] = src
	}
}

func jnn(vm *VM, opr uint8) {
	var src uint32

	if opr == 0 {
		vm.fetch(4)
		src = vm.reg[ir]
	} else {
		vm.fetch(1)
		r := uint8(vm.reg[ir])
		src = vm.reg[r]
	}

	nfs := ((1 << nf) & vm.reg[flg]) >> nf
	if nfs == 0 {
		vm.reg[pc] = src
	}
}

var opLookup = [64]opFn{
	nop,
	mov8,
	mov16,
	mov32,
	add,
	sub,
	//mul
	//div
	and,
	or,
	not,
	xor,
	jmp,
	jz,
	jnz,
	jn,
	jnn,
}