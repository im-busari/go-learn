package mocking

import "time"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration) // signature is the same as time.Sleep
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}
