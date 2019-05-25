package slices

func Sum(numbers []int) int {
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

func SumAll(setOfNumbers ...[]int) []int {

	// 73.9 ns/op
	var result []int

	for _, numbers := range setOfNumbers {
		result = append(result, Sum(numbers))
	}

	/*
		// 42.5 ns/op
		result := make([]int, len(setOfNumbers))
		for i, numbers := range setOfNumbers {
			result[i] = Sum(numbers)
		}
	*/

	return result
}

func SumAllTails(setOfNumbers ...[]int) []int {
	var result []int
	//  numbers[len(numbers)-1]
	for _, numbers := range setOfNumbers {
		if len(numbers) == 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, numbers[len(numbers)-1])
	}

	return result
}
