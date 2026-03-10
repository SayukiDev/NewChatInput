package aivis

import (
	execCmd "ChatInput/pkg/cmd"
	"ChatInput/pkg/portkill"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type Engine struct {
	f        *os.File
	exitChan chan error
	*exec.Cmd
}

func NewEngine(logPath, execPath string, args ...string) (*Engine, error) {
	c := exec.Command(execPath, args...)
	lofFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.SysProcAttr = execCmd.HideWindowAttr()
	c.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8", "PYTHONUNBUFFERED=1")
	c.Stdout = lofFile
	c.Stderr = lofFile
	return &Engine{
		f:        lofFile,
		exitChan: make(chan error, 1),
		Cmd:      c,
	}, nil
}

func (e *Engine) IsRunning() (bool, error) {
	if e.Cmd.Process == nil {
		return false, nil
	}
	select {
	case err := <-e.exitChan:
		return false, fmt.Errorf("the process exited: %w", err)
	default:
	}
	return true, nil
}

func (e *Engine) InstallModel(modelPath string) error {
	modelPath = strings.ToLower(modelPath)
	if !strings.Contains(path.Ext(modelPath), ".aivmx|.aivm") {
		return fmt.Errorf("invalid model file: %s", modelPath)
	}
	f, err := os.Open(modelPath)
	if err != nil {
		return err
	}
	defer f.Close()
	ucp, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	f2, err := os.OpenFile(path.Join(ucp, "AivisSpeech-Engine\\Models"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f2.Close()
	_, err = io.Copy(f2, f)
	return nil
}

func (e *Engine) Start() error {
	port := 10101
	for i, a := range e.Args {
		if strings.TrimSpace(a) != "--port" {
			continue
		}
		if i+1 >= len(e.Args) {
			break
		}
		p, err := strconv.Atoi(e.Args[i+1])
		if err != nil {
			return fmt.Errorf("invalid port: %w", err)
		}
		port = p
	}
	portkill.KillPort(strconv.Itoa(port))
	err := e.Cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		e.exitChan <- e.Wait()
		close(e.exitChan)
	}()
	return nil
}

func (e *Engine) Close() error {
	defer e.f.Close()
	if e.Cmd.Process == nil {
		return nil
	}
	e.Cmd.Process.Kill()
	select {
	case <-e.exitChan:
	}
	return nil
}
