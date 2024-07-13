package Chapter_4

// Sum get the sum of integers in a slices returns sum total.
func Sum(numbers []int) int {
	var total int
	for _, itr := range numbers {
		total += itr
	}

	return total
}

// SumAll adds up the values in the slices and returns a new slice with the summed values.
func SumAll(slices ...[]int) []int {
	lengthOfNumbers := len(slices)
	sums := make([]int, lengthOfNumbers)

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) < 2 {
			continue
		}

		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
