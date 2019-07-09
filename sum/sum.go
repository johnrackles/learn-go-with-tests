package sum

// Sum calculates the sum from an array of numbers
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// All calculates the sum of all sums of all passed slices
func All(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// AllTails calcultes the sum of all tails
func AllTails(tailsToSum ...[]int) (sums []int) {
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
