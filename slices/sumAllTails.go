package slices

func SumAllTails(listOfNumbers ...[]int) []int {
	sum := []int{}

	for _, numbers := range listOfNumbers {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}