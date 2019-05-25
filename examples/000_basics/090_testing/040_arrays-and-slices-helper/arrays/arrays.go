package arrays

func Sum(numbers [5]int) int {
	sum := 0
	/*
		count := len(numbers)
		for i := 0; i < count; i++ {
			sum += numbers[i]
		}
	*/

	for _, number := range numbers {
		sum += number

	}

	return sum
}
