package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger() {
	// Ensure logs directory exists
	logDir := "logs"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create log directory:", err)
		os.Exit(1)
	}

	// Create log file with date
	today := time.Now().Format("2006-01-02")
	logPath := filepath.Join(logDir, today+".txt")

	LogFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}

	Logger = log.New(LogFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(LogFile) // Redirect standard log output
}

// Optional: close file when app exits
func CloseLogger() {
	if LogFile != nil {
		LogFile.Close()
	}
}
