package mocking

import (
	"fmt"
	"io"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

// or could also use (out *bytes.Buffer)
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(out, finalWord)
}
