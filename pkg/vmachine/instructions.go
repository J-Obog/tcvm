package vmachine

func (vm *VM) getSrc(stype uint8, size uint8) uint32 {
	switch stype {
	case M_REG: // register
		reg := vm.ram[vm.pc]
		vm.pc++
		return vm.regRead(reg, size)

	case M_EREG: // [register]
		reg := vm.ram[vm.pc]
		vm.pc++
		addr := vm.regs[reg]
		return vm.memRead(addr, size)

	case M_EMEM: // [memory]
		addr := vm.memRead(vm.pc, DWORD)
		vm.pc += 4
		return vm.memRead(addr, size)

	case M_IMMED: // immediate
		v := vm.memRead(vm.pc, size)
		vm.pc += uint32(size)
		return v
	}

	return 0
}

func (vm *VM) RSType() (size uint8, source uint8, register uint8) {
	hdr := vm.ram[vm.pc]
	vm.pc++

	sz := (hdr >> 5) & 0x7
	src := (hdr >> 3) & 0x3
	reg := hdr & 0x7

	return sz, src, reg
}

func (vm *VM) MType() (size uint8, source uint8, destination uint8) {
	hdr := vm.ram[vm.pc]
	vm.pc++

	sz := (hdr >> 4) & 0x7
	src := (hdr >> 2) & 0x3
	dest := hdr & 0x3

	return sz, src, dest
}

func (vm *VM) JType() (negation uint8, source uint8, flag uint8) {
	hdr := vm.ram[vm.pc]
	vm.pc++

	neg := (hdr >> 5) & 0x1
	src := (hdr >> 3) & 0x3
	flg := hdr & 0x7

	return neg, src, flg
}

type opFn func(vm *VM)

func nop(vm *VM) {
	//no operation

}

func mov(vm *VM) {

}

func add(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] += vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func sub(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] += (^vm.getSrc(s, sz))
	vm.updateFlags(vm.regs[r])
}

func mul(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] *= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func div(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] /= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func and(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] &= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func or(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] |= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func not(vm *VM) {
	_, _, r := vm.RSType()

	vm.regs[r] = ^vm.regs[r]
	vm.updateFlags(vm.regs[r])
}

func xor(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] ^= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func cmp(vm *VM) {
	sz, s, r := vm.RSType()

	tmp := vm.regs[r] + (^vm.getSrc(s, sz))
	vm.updateFlags(tmp)
}

func jmp(vm *VM) {

}

func push(vm *VM) {

}

func pop(vm *VM) {

}

func call(vm *VM) {

}

func ret(vm *VM) {

}

func shl(vm *VM) {

}

func shr(vm *VM) {

}

func sys(vm *VM) {

}

var opLookup = [32]opFn{
	nop,
	/*mov,
	add,
	sub,
	mul,
	div,
	and,
	or,
	not,
	xor,
	cmp,
	jmp,
	push,
	pop,
	call,
	ret,
	shl,
	shr,
	sys,*/
}
