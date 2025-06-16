//go:build !windows

package cmd

import "syscall"

func HideWindowAttr() *syscall.SysProcAttr {
	return nil
}
