package integers

func BubbleSort(nums []int) []int {
	for i := range nums {
		swapped := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return nums
}

func swap(nums []int, position int) {
	aux := nums[position]
	nums[position] = nums[position+1]
	nums[position+1] = aux
}
