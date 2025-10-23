package arrays

// An interesting property of arrays is that the size is encoded in its type
func SumArr(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumSlice(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	// make allows you to create a slice with a starting capacity
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		// sums = append(sums, SumSlice(numbers)) --> not concerned with capacity
		sums[i] = SumSlice(numbers)
	}

	return sums
}

// Calculates the totals of the "tails" of each slice. The tail of a collection is all items in the collection
// except the first one (the "head").
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		
		tail := numbers[1:]
		sums = append(sums, SumSlice(tail))
	}

	return sums
}
