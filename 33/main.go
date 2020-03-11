package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 1}, 2))
	fmt.Println(search([]int{5, 1, 3}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2, 3}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	s, e, m := 0, len(nums)-1, 0
	for s <= e {
		m = (s + e) / 2
		if target == nums[m] {
			return m
		}
		if nums[s] <= nums[m] {
			if nums[m] > target && nums[s] <= target {
				e = m - 1
			} else {
				s = m + 1
			}
		} else {
			if nums[m] < target && nums[e] >= target {
				s = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}
