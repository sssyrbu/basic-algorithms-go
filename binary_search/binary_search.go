package main

import "fmt"

// function returns a number of the given target in the given list
func binary_search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		switch {
		case target > guess:
			low = mid + 1
		case target < guess:
			high = mid - 1
		default: // target == guess:
			return mid
		}
	}
	return -1 // not found
}

func main() {
	list := []int{1, 3, 7, 9, 14, 25}
	fmt.Println(binary_search(list, 1))
}
