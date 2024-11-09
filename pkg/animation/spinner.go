package spinner

import (
	"log"
	"time"
)

type Spinner struct {
	chars  []rune
	delay  time.Duration
	active bool
	stop   chan bool
	suffix string
}

func New(chars []rune, delay time.Duration) *Spinner {
	return &Spinner{
		chars: chars,
		delay: delay,
		stop:  make(chan bool),
	}
}

func (s *Spinner) Start() {
	s.active = true
	go func() {
		for {
			for _, r := range s.chars {
				select {
				case <-s.stop:
					return
				default:
					log.Printf("\r%c %s", r, s.suffix)
					time.Sleep(s.delay)
				}
			}
		}
	}()
}

func (s *Spinner) Stop() {
	if s.active {
		s.stop <- true
		s.active = false
		log.Print("\r   \r") // Clear the spinner
	}
}

func (s *Spinner) Suffix(text string) {
	s.suffix = text
}
