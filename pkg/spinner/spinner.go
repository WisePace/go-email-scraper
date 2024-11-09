package spinner

import (
	"fmt"
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

func (spinner *Spinner) Start() {
	spinner.active = true
	go func() {
		for {
			for _, r := range spinner.chars {
				select {
				case <-spinner.stop:
					return
				default:
					fmt.Printf("\r%c %s", r, spinner.suffix)
					time.Sleep(spinner.delay)
				}
			}
		}
	}()
}

func (spinner *Spinner) Stop() {
	if spinner.active {
		spinner.stop <- true
		spinner.active = false
		fmt.Print("\r   \r") // Clear the spinner
	}
}

func (spinner *Spinner) Suffix(text string) {
	spinner.suffix = text
}
