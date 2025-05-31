package daemon

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WritePID(pid int) error {
	logFile, err := os.OpenFile(PID_FILE, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer logFile.Close()
	fmt.Fprintf(logFile, "%d\n", pid)
	return nil
}

func ReadPID() (int, error) {
	data, err := os.ReadFile(PID_FILE)
	if err != nil {
		return -1, err
	}
	pidStr := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return -1, fmt.Errorf("invalid PID format: %v", err)
	}

	return pid, nil
}

func DeletePID() error {
	err := os.Remove(PID_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete PID file: %v", err)
	}
	return nil
}
