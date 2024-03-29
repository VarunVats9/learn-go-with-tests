package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		countdownOperationsSpy := &CountdownOperationsSpy{}
		Countdown(buffer, countdownOperationsSpy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before very print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
