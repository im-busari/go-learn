package mocking

import "iter"

// Countdown using iterators: https://tip.golang.org/doc/go1.23
func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
