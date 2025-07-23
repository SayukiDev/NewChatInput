package portkill

import (
	"ChatInput/pkg/cmd"
	"fmt"
	"os/exec"
	"runtime"
)

func KillPort(port string) error {
	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd := fmt.Sprintf("Stop-Process -Id (Get-NetTCPConnection -LocalPort %s).OwningProcess -Force", port)
		command = exec.Command("powershell.exe", "-Command", cmd)
	} else {
		cmd := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		command = exec.Command("bash", "-c", cmd)
	}
	err := execCmd(command)
	if err != nil {
		return fmt.Errorf("failed to kill port %s: %w", port, err)
	}
	return nil
}

func execCmd(cmd *exec.Cmd) error {
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}
	return nil
}

func IsPortOpen(port string) bool {
	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		cmdS := fmt.Sprintf("Get-NetTCPConnection -LocalPort %s", port)
		command = exec.Command("powershell.exe", "-Command", cmdS, "-WindowStyle", "Hidden")
	} else {
		cmdS := fmt.Sprintf("lsof -i tcp:%s", port)
		command = exec.Command("bash", "-c", cmdS)
	}
	command.SysProcAttr = cmd.HideWindowAttr()
	err := execCmd(command)
	if err != nil {
		return false
	}
	return true
}
