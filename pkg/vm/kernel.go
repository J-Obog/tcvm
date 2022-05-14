package vm

//sys call mappings
const (
	SYS_EXIT  uint32 = 0
	SYS_READ  uint32 = 1
	SYS_WRITE uint32 = 2
	SYS_OPEN  uint32 = 3
	SYS_CLOSE uint32 = 4
	SYS_SBRK  uint32 = 5
)


//system call 
func (c *Cpu) sysCall(num uint32) {
	switch num {

	}
}