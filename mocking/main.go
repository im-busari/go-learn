package mocking

import (
	"os"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	
	Countdown(os.Stdout, sleeper)

	//Countdown(os.Stdout, &DefaultSleeper{})
}
