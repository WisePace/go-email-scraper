package log

import (
	"os"
	"testing"
)

func TestOpenFile(testCase *testing.T) {
	fileName := "test.log"
	file, err := OpenFile(fileName)
	if err != nil {
		testCase.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()
	defer os.Remove(fileName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		testCase.Errorf("Expected file %s to exist", fileName)
	}
}
