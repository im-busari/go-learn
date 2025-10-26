package mocking

import (
	"fmt"
	"io"
	"time"
)

// Shared interface with the DefaultSleeper and SpySleeper which is being used in the tests
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const countdownStart = 3
const finalWord = "Go!"

// or could also use (out *bytes.Buffer)
func Countdown(out io.Writer, sleeper Sleeper) {
	//for i := countdownStart; i > 0; i-- {
	//	sleeper.Sleep()
	//}
	//
	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	//}
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
