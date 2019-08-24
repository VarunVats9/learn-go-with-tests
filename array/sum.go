package array

// Sum takes array as input, and returns the sum of all of them.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll compute sum all the given slices, and return this in new slice.
func SumAll(numbersToSum ...[]int) (sum []int) {
	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}
	return
}

// SumAllTails compute sum all the given slices except the head element, and return this in new slice.
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}
	return
}
