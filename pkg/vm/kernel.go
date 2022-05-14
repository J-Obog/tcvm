package vm

import (
	"os"

	"github.com/J-Obog/tcvm/pkg/com"
)

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
	case SYS_EXIT:
		exitCode := c.Registers[com.R1]
		os.Exit(int(exitCode))
	
	
	case SYS_SBRK:
		incr := c.Registers[com.R1]
		oldBrk := c.ESP
		newBrk := c.ESP - incr

		if (newBrk <= c.SBP) || (newBrk <= c.Registers[com.SP]) {
			c.Registers[com.R0] = 0xFFFFFFFF
		} else {
			c.ESP = newBrk
			c.Registers[com.R0] = oldBrk
		}	
	}
}
