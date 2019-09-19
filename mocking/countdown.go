package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper interfac
type Sleeper interface {
	Sleep()
}

// SpySleeper contains Calls field
type SpySleeper struct {
	Calls int
}

// Sleep counts number of times it is called
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Countdown prints 3, 2, 1, Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
