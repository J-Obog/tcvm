package vmachine

import "fmt"

type opFn func(*VM, uint8)

func operands(vm *VM, isOp2 bool, mode uint8, opsz uint8) (uint8, uint32) { // convenience method for getting operands
	if isOp2 {
		vm.fetch(1)
	}

	d := uint8(vm.reg[ir])

	if mode%2 != 0 {
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
	dest, src := operands(vm, true, mode, 1)

	if mode <= 1 {
		vm.reg[dest] = uint32(uint8(src))
	} else {
		vm.mem[vm.reg[dest]] = uint8(src)
	}
}

func mov16(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 2)

	if mode <= 1 {
		vm.reg[dest] = uint32(uint16(src))
	} else {
		vm.write(vm.reg[dest], 2, src)
	}
}

func mov32(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)

	if mode <= 1 {
		vm.reg[dest] = src
	} else {
		vm.write(vm.reg[dest], 4, src)
	}
}

func add(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] = uint32(uint64(vm.reg[dest]) + uint64(src))
	vm.updateFlags(dest)
}

func sub(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] = uint32(uint64(vm.reg[dest]) + uint64((^src)+1))
	vm.updateFlags(dest)
}

func and(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] &= src
	vm.updateFlags(dest)
}

func or(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] |= src
	vm.updateFlags(dest)
}

func xor(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] ^= src
	vm.updateFlags(dest)
}

func not(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] = ^src
	vm.updateFlags(dest)
}

func jmp(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	vm.reg[pc] = src
}

func jz(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	if vm.checkFlag(zf) {
		vm.reg[pc] = src
	}
}

func jnz(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	if !vm.checkFlag(zf) {
		vm.reg[pc] = src
	}
}

func jn(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	if vm.checkFlag(nf) {
		vm.reg[pc] = src
	}
}

func jnn(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	if !vm.checkFlag(nf) {
		vm.reg[pc] = src
	}
}

func push8(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 1)
	vm.mem[vm.reg[sp]] = uint8(src)
	vm.reg[sp]++
}

func push16(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 2)
	vm.write(vm.reg[sp], 2, src)
	vm.reg[sp] += 2
}

func push32(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	vm.write(vm.reg[sp], 4, src)
	vm.reg[sp] += 4
}

func pop8(vm *VM, mode uint8) {
	vm.reg[sp]--
}

func pop16(vm *VM, mode uint8) {
	vm.reg[sp] -= 2
}

func pop32(vm *VM, mode uint8) {
	vm.reg[sp] -= 4
}

func call(vm *VM, mode uint8) {
	_, src := operands(vm, false, mode, 4)
	vm.reg[rar] = vm.reg[pc] + 1
	vm.reg[pc] = src
}

func ret(vm *VM, mode uint8) {
	vm.reg[pc] = vm.reg[rar]
}

func shl(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] <<= src
	vm.updateFlags(dest)
}

func shr(vm *VM, mode uint8) {
	dest, src := operands(vm, true, mode, 4)
	vm.reg[dest] >>= src
	vm.updateFlags(dest)
}

func sys(vm *VM, mode uint8) {
	num := uint8(vm.reg[r5])

	switch num {
		case halt:
			vm.mem[flg] |= (1 << hf)
			fmt.Print("Program exited")
		break

		case puts:
			ptr := vm.reg[r6]

			for vm.mem[ptr] != 0 {
				fmt.Print(string(vm.mem[ptr]))
				ptr++
			}
		break
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
	push8,
	push16,
	push32,
	pop8,
	pop16,
	pop32,
	call,
	ret,
	shl,
	shr,
	sys,
}