package main

import (
	"log"
	"os"
)

// OpenLogFile opens a log file and appends the output of scraper
// to the file. It returns the file and a logger.
func OpenLogFile(fileName string) (*os.File, *log.Logger) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening %s: %v", fileName, err)
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return logFile, logger
}
