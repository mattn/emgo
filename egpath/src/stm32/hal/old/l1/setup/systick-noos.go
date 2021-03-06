// build +noos

package setup

import (
	"runtime/noos/sysclk/cmst"
	"syscall"
)

func sysclkChanged() {
	if syscall.MaxTasks() == 0 {
		return
	}
	lev, _ := syscall.SetPrivLevel(0)
	cmst.Setup(2e6, SysClk, false)
	syscall.SetSysClock(cmst.Uptime, cmst.SetWakeup)
	syscall.SetPrivLevel(lev)
}
