package vmachine

import "fmt"


func (vm *VM) getDest(dtype uint8) uint32 {
	switch dtype {
	case M_REG: // register
		reg := vm.ram[vm.pc]
		vm.pc++
		return uint32(reg)

	case M_EREG: // [register]
		reg := vm.ram[vm.pc]
		vm.pc++
		addr := vm.regs[reg]
		return addr

	case M_EMEM: // [memory]
		addr := vm.memRead(vm.pc, DWORD)
		vm.pc += 4
		return addr
	}

	return 0
}

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
	sz, s, d := vm.MType()
	loc := vm.getDest(d)
	val := vm.getSrc(s, sz)

	if d == M_REG {
		vm.regWrite(uint8(loc), sz, val)
	} else {
		vm.memWrite(loc, sz, val)
	}
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
	n, s, f := vm.JType()

	addr := vm.getSrc(s, DWORD)

	c := vm.getFlag(f)

	if n == 1 {
		c = !c
	}

	if c {
		vm.pc = addr
	}
}

func push(vm *VM) {
	sz, s, _ := vm.RSType()

	addr := vm.regs[R_SP]
	data := vm.getSrc(s, sz)

	vm.memWrite(addr, sz, data)
	vm.regs[R_SP] += uint32(sz)
}

func pop(vm *VM) {
	sz, _, _ := vm.RSType()

	vm.regs[R_SP] -= uint32(sz)
}

func call(vm *VM) {
	sz, s, _ := vm.RSType()

	vm.rar = vm.pc + 1
	vm.pc = vm.getSrc(s, sz)
}

func ret(vm *VM) {
	vm.pc = vm.rar
}

func shl(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] <<= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func shr(vm *VM) {
	sz, s, r := vm.RSType()

	vm.regs[r] >>= vm.getSrc(s, sz)
	vm.updateFlags(vm.regs[r])
}

func sys(vm *VM) {
	syscall := uint8(vm.regs[R_R5])
	switch syscall {
	case SYS_HALT:
		vm.flags |= (1 << F_HALT)
		fmt.Print("Program exited")

	case SYS_PUTS:
		ptr := vm.regs[R_R6]
		for vm.ram[ptr] != 0 {
			fmt.Print(string(vm.ram[ptr]))
			ptr++
		}
		
	case SYS_GETS:
		ptr := vm.regs[R_R6]
		var buf []byte
		fmt.Scanln(&buf)
		for _, c := range buf {
			vm.ram[ptr] = c
			ptr++
		}
		vm.ram[ptr] = 0
	}
}

var opLookup = [32]opFn{
	nop,
	mov,
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
	sys,
}
