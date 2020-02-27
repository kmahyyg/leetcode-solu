package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{3,0,3,2,-4,0,-3,2,2,0,-1,-5}))
}

func threeSum(nums []int) [][]int {
	// sort data samples ascending first
	sort.Ints(nums)
	// define a starting number
	var solus = make([][]int, 0)
	for i:=0; i < len(nums) - 2 ; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue    // prevent repeat number from being calculated again
		}
		start := i+1
		end := len(nums) - 1
		target := -nums[i]
		for start < end {
			sum := nums[start] + nums[end]
			switch {
			case sum < target:
				start++
			case sum > target:
				end--
			default:
				solus = append(solus, []int{nums[i], nums[start], nums[end]})
				for start + 1 < end && nums[start] == nums[start + 1] {
					start++  // move forward to the next same number (mid)
				}
				for end - 1 > start && nums[end] == nums[end - 1] {
					end--   // move back to the next same number (last)
				}
				start++ // continue, prevent deadloop
				end--  // continue, prevent deadloop
			}
		}
	}
	return solus
}
