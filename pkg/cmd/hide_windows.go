// build
package cmd

import "syscall"

func HideWindowAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow: true,
	}
}
