package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 1}, 2))
	fmt.Println(search([]int{5, 1, 3}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

// binary search here
// if nums[left:mid] < nums[mid] < nums[mid:right], then
// nums[left:mid] must be increasing only
func search(nums []int, target int) int {
	current, mid, last := 0, 0, len(nums)
	for current != last {
		mid = current + (last - current) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[current] <= nums[mid] {
			if nums[current] <= target && target < nums[mid] {
				last = mid
			} else {
				current = mid + 1
			}
		} else {
			if nums[current] < target && target <= nums[last-1] {
				current = mid + 1
			} else {
				last = mid
			}
		}
	}
	return -1
}
