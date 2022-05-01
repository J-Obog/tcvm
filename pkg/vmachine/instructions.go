package vmachine

func (vm *VM) getSrc(stype uint8, size uint8) uint32 {
	switch stype {
	case M_REG:
		reg := vm.memRead(vm.pc, BYTE)
		vm.pc++
		return vm.regRead(reg, size)

	case M_EREG:
		reg := vm.memRead(vm.pc, BYTE)
		vm.pc++
		addr := vm.regRead(reg, DWORD)
		return vm.memRead(addr, size)

	case M_EMEM:
		addr := vm.memRead(vm.pc, DWORD)
		vm.pc += 4
		return vm.memRead(addr, size)

	case M_IMMED:
		v := vm.memRead(vm.pc, size)
		vm.pc += uint32(size)
		return v
	}

	return 0
}

func (vm *VM) RSType() {
	//register-source
}

func (vm *VM) MType() {
	//mov
}

func (vm *VM) JType() {
	//jmp
}

type opFn func(vm *VM)

func nop(vm *VM) {
	//no operation

}

func mov(vm *VM) {

}

func add(vm *VM) {

}

func sub(vm *VM) {

}

func mul(vm *VM) {

}

func div(vm *VM) {

}

func and(vm *VM) {

}

func or(vm *VM) {

}

func not(vm *VM) {

}

func xor(vm *VM) {

}

func cmp(vm *VM) {

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
