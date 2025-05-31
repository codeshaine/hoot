package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/codeshaine/hoot/internal/logging"
	"github.com/codeshaine/hoot/internal/notify"
)

type ProcessStatus string

const (
	RUNNING ProcessStatus = "RUNNING"
	IDLE    ProcessStatus = "IDLE"
)

const PID_FILE = "/tmp/hoot.pid"

func StartDaemon(interval time.Duration) error {
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}
	cmd := exec.Command(executable, "start","--daemon","--interval",interval.String())

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	logFile, err := os.OpenFile(logging.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer logFile.Close()

	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Stdin = nil

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start daemon: %v", err)
	}

	pid := cmd.Process.Pid
	err = WritePID(pid)
	if err != nil {
		return err
	}
	return nil
}

func StartHoot(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		err := notify.PlayNotificationSound(nil)
		if err != nil {
			logging.Log(err.Error())
		}
		err = notify.PushDesktopNotification(nil)
		if err != nil {
			logging.Log(err.Error())
		}
	}
}

func StopHoot() error {
	pid, err := ReadPID()
	if err != nil {
		return err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = process.Kill()
	if err != nil {
		return err
	}
	err = DeletePID()
	if err != nil {
		return err
	}
	return nil
}

func Status() ProcessStatus {
	pid, err := ReadPID()
	if err != nil {
		return IDLE
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return IDLE
	}

	// On Unix-like systems, sending signal 0 checks existence
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return IDLE // process doesn't exist
	}

	return RUNNING
}
