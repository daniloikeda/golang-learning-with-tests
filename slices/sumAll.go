package slices 

func SumAll(listOfNumbers ...[]int) []int {
	sumAll := []int{}
	for _, numbers := range listOfNumbers {
		sumAll = append(sumAll, Sum(numbers))
	}
	
	return sumAll
}