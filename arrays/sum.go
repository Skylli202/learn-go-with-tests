package arrays

func Sum(numbers []int) int {
	return Reduce(numbers, func(acc, x int) int { return acc + x }, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	res := []int{}
	return Reduce(numbersToSum, func(acc, x []int) []int { return append(acc, Sum(x)) }, res)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	result := initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
