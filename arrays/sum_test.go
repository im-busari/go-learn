package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArr(numbers)
		want := 15

		assertCorrectValue(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		assertCorrectValue(t, got, want)
	})

	t.Run("sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertCorrectSlices(t, got, want)
	})

	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertCorrectSlices(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 4, 7}, []int{})
		want := []int{13, 0}

		assertCorrectSlices(t, got, want)
	})
}

func assertCorrectValue(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func assertCorrectSlices(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) { // or we could also use !reflect.DeepEqual(got, want)
		// %v to print an array/slice
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
