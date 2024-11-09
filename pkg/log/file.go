package log

import (
	"os"
)

// OpenFile opens a log file and appends the output of scraper
// to the file. It returns the file.
func OpenFile(fileName string) (*os.File, error) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
