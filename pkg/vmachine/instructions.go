package vmachine

type opFn func(*VM, uint8)

func operands(vm *VM, isOp2 bool, isReg bool, opsz uint8) (uint8, uint32) { // convenience method for getting operands
	if isOp2 {
		vm.fetch(1)
	}

	d := uint8(vm.reg[ir])

	if isReg {
		vm.fetch(1)
		return d, vm.reg[uint8(vm.reg[ir])]
	}

	vm.fetch(opsz)
	return d, vm.reg[ir]
}

func nop(vm *VM, mode uint8) {
	//no operation
}

func mov8(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 1)

	if mode <= 1 {
		vm.reg[dest] = uint32(uint8(src))
	} else {
		vm.mem[vm.reg[dest]] = uint8(src)
	}
}

func mov16(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 2)

	if mode <= 1 {
		vm.reg[dest] = uint32(uint16(src))
	} else {
		vm.mem[vm.reg[dest]] = uint8(src >> 8)
		vm.mem[vm.reg[dest]+1] = uint8(src)
	}
}

func mov32(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)

	if mode <= 1 {
		vm.reg[dest] = src
	} else {
		vm.mem[vm.reg[dest]] = uint8(src >> 24)
		vm.mem[vm.reg[dest]+1] = uint8(src >> 16)
		vm.mem[vm.reg[dest]+2] = uint8(src >> 8)
		vm.mem[vm.reg[dest]+3] = uint8(src)
	}
}

func add(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] = uint32(uint64(vm.reg[dest]) + uint64(src))
	vm.setFlags(dest)
}

func sub(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] = uint32(uint64(vm.reg[dest]) + uint64((^src)+1))
	vm.setFlags(dest)
}

func and(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] &= src
	vm.setFlags(dest)
}

func or(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] |= src
	vm.setFlags(dest)
}

func xor(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] ^= src
	vm.setFlags(dest)
}

func not(vm *VM, mode uint8) {
	dest, src := operands(vm, true, (mode%2 != 0), 4)
	vm.reg[dest] = ^src
	vm.setFlags(dest)
}

func jmp(vm *VM, mode uint8) {
	_, src := operands(vm, false, (mode%2 != 0), 4)
	vm.reg[pc] = src
}

func jz(vm *VM, mode uint8) {
	_, src := operands(vm, false, (mode%2 != 0), 4)
	if vm.checkFlag(zf) {
		vm.reg[pc] = src
	}
}

func jnz(vm *VM, mode uint8) {
	_, src := operands(vm, false, (mode%2 != 0), 4)
	if !vm.checkFlag(zf) {
		vm.reg[pc] = src
	}
}

func jn(vm *VM, mode uint8) {
	_, src := operands(vm, false, (mode%2 != 0), 4)
	if vm.checkFlag(nf) {
		vm.reg[pc] = src
	}
}

func jnn(vm *VM, mode uint8) {
	_, src := operands(vm, false, (mode%2 != 0), 4)
	if !vm.checkFlag(nf) {
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