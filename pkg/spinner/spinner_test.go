package spinner

import (
	"testing"
	"time"
)

func TestSpinnerStartStop(testCase *testing.T) {
	spinner := New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	spinner.Start()
	time.Sleep(500 * time.Millisecond)
	spinner.Stop()
	if spinner.active {
		testCase.Error("Expected spinner to be inactive after Stop()")
	}
}

func TestSpinnerSuffix(testCase *testing.T) {
	spinner := New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	spinner.Suffix("test suffix")
	if spinner.suffix != "test suffix" {
		testCase.Errorf("Expected suffix to be 'test suffix', got '%s'", spinner.suffix)
	}
}
