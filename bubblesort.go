package main

import (
	"fmt"
	"os"
)

func fillSlice(size int) []int {
	data := make([]int, 0, size)
	fmt.Printf("Enter %d values for the slice.\n", size)
	for i := 0; i < size; i++ {
		var val int
		fmt.Printf("[%2d/%d]: ", i+1, size)
		_, err := fmt.Scan(&val)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		data = append(data, val)
	}
	return data
}

func Swap(data []int, i int) {
	tmp := data[i]
	data[i] = data[i+1]
	data[i+1] = tmp
}

func BubbleSort(data []int) {
	for n := len(data); n > 1; {
		newN := 0
		for i := 1; i < n; i++ {
			if data[i-1] > data[i] {
				Swap(data, i-1)
				newN = i
			}
		}
		n = newN
	}
}

func main() {
	fmt.Print("How many values to enter?: ")
	var len int
	_, err := fmt.Scan(&len)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	data := fillSlice(len)
	BubbleSort(data)
	fmt.Print(data)
}
