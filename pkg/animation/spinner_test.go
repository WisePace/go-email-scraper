package spinner

import (
	"testing"
	"time"
)

func TestSpinnerStartStop(t *testing.T) {
	s := New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	s.Start()
	time.Sleep(500 * time.Millisecond)
	s.Stop()
	if s.active {
		t.Error("Expected spinner to be inactive after Stop()")
	}
}

func TestSpinnerSuffix(t *testing.T) {
	s := New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	s.Suffix("test suffix")
	if s.suffix != "test suffix" {
		t.Errorf("Expected suffix to be 'test suffix', got '%s'", s.suffix)
	}
}
