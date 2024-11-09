package log

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	fileName := "test.log"
	file, err := OpenFile(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()
	defer os.Remove(fileName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist", fileName)
	}
}
