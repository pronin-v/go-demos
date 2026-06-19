package arraysnslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}
	return result
}
