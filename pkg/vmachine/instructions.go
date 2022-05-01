package vmachine

import "reflect"

func (vm *VM) getDest(dtype uint8) (dest Store, destLoc uint32) {
	switch dtype {
	case M_REG:
		reg := vm.ram.Read(vm.pc, BYTE)
		vm.pc++
		return &vm.registers, reg

	case M_EREG:
		reg := vm.ram.Read(vm.pc, BYTE)
		vm.pc++
		addr := vm.registers.Read(reg, DWORD)
		return &vm.ram, addr

	case M_EMEM:
		addr := vm.ram.Read(vm.pc, DWORD)
		vm.pc += 4
		return &vm.ram, addr
	}

	return nil, 0
}

func (vm *VM) getSrc(stype uint8, size uint8) uint32 {
	switch stype {
	case M_REG:
		reg := vm.ram.Read(vm.pc, BYTE)
		vm.pc++
		return vm.registers.Read(reg, size)

	case M_EREG:
		reg := vm.ram.Read(vm.pc, BYTE)
		vm.pc++
		addr := vm.registers.Read(reg, DWORD)
		return vm.ram.Read(addr, size)

	case M_EMEM:
		addr := vm.ram.Read(vm.pc, DWORD)
		vm.pc += 4
		return vm.ram.Read(addr, size)

	case M_IMMED:
		v := vm.ram.Read(vm.pc, size)
		vm.pc += uint32(size)
		return v
	}

	return 0
}

func (vm *VM) getDSType() (dest *Store, destLoc uint32, src uint32, dataSize uint8) {
	b := uint8(vm.ram.Read(vm.pc, BYTE))
	vm.pc++
	dt := (b >> 2) & 0x3
	st := b & 0x3
	sz := (b >> 4) & 0xF

	d, dl := vm.getDest(dt)
	s := vm.getSrc(st, sz)

	return &d, dl, s, sz

}

func (vm *VM) getDType() (dest *Store, destLoc uint32, dataSize uint8) {
	b := uint8(vm.ram.Read(vm.pc, BYTE))
	vm.pc++
	dt := b & 0x3
	sz := (b >> 2) & 0xF

	d, dl := vm.getDest(dt)

	return &d, dl, sz
}

func (vm *VM) getSType() (src uint32, dataSize uint8) {
	b := uint8(vm.ram.Read(vm.pc, BYTE))
	vm.pc++
	st := b & 0x3
	sz := (b >> 2) & 0xF

	s := vm.getSrc(st, sz)

	return s, sz
}

func (vm *VM) getJType() (neg uint8, flag uint8, src uint32) {
	b := uint8(vm.ram.Read(vm.pc, BYTE))
	vm.pc++
	st := b & 0x3
	flg := (b >> 2) & 0xF
	ng := (b >> 6) & 0x1

	s := vm.getSrc(st, 0x4)

	return ng, flg, s
}

type opFn func(vm *VM)

func nop(vm *VM) {
	//no operation

}

func mov(vm *VM) {
	d, dl, s, sz := vm.getDSType()
	va := reflect.ValueOf(d).Elem()
	va.Write()
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
