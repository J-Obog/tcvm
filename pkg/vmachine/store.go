package vmachine

type Store interface {
	Read(loc uint32, rsize uint8) uint32
	Write(loc uint32, wsize uint8, data uint32)
}

type Memory struct {
	Store
	mem [MAX_MEM_SIZE]uint8
}

func (m *Memory) Read(loc uint32, rsize uint8) uint32 {
	ptr := loc
	end := ptr + uint32(rsize)
	data := uint32(0)

	for ptr < end {
		data <<= 8
		data |= uint32(m.mem[ptr])
		ptr++
	}

	return data
}

func (m *Memory) Write(loc uint32, wsize uint8, data uint32) {
	ptr := loc
	end := ptr + uint32(wsize)
	sfac := wsize - 1

	for ptr < end {
		word := (data & (255 << (8 * sfac))) >> (8 * sfac)
		m.mem[ptr] = uint8(word)
		ptr++
		sfac--
	}
}

type RegisterFile struct {
	Store
	regs [8]uint32
}

func (r *RegisterFile) Read(loc uint32, rsize uint8) uint32 {
	return r.regs[loc] & ((1 << (8 * rsize)) - 1)
}

func (r *RegisterFile) Write(loc uint32, wsize uint8, data uint32) {
	d := (data & ((1 << (8 * wsize)) - 1))
	r.regs[loc] = d
}