package logging

import (
	"fmt"
	"os"
	"time"
)

const LOG_FILE = "/tmp/hoot.log"

func Log(message string) {
	logFile, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer logFile.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(logFile, "[%s] %s\n", timestamp, message)
}
