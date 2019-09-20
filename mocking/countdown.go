package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// CountdownOperationsSpy will have list of operations
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep adds the sleep operation to the list of operations
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write adds the write operation to the list of operations
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
