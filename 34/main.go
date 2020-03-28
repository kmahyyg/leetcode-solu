package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{3, 3, 3}, 3))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange_v1(nums []int, target int) []int {
	// brute force
	start := -1
	found := 0
	res := []int{}
	for k, v := range nums {
		if v == target {
			found = 1
			start = k
			res = append(res, start)
			break
		}
	}
	if start == -1 {
		return []int{-1, -1}
	}
	last := start
	for i := start + 1; i < len(nums); i++ {
		if nums[i] == target {
			last = i
			found += 1
		}
	}
	if found == 1 {
		return append(res, res[0])
	}
	res = append(res, last)
	return res
	// Runtime: 8 ms, faster than 83.02% of Go online submissions for
	//     Find First and Last Position of Element in Sorted Array.
	// Memory Usage: 4.1 MB, less than 50.00% of Go online submissions for
	//     Find First and Last Position of Element in Sorted Array.
}

func searchHelper(nums []int, target int, leftmost bool) int {
	lo := 0
	hi := len(nums) // v2_nohelper: len(nums)-1
	for lo < hi {
		idx := lo + (hi-lo)/2
		// find leftmost
		if nums[idx] > target || (leftmost && nums[idx] == target) {
			hi = idx
		} else { // v2_nohelper : } else if nums[idx] < target {
			lo = idx + 1
		}
	}
	return lo
}

func searchRange_v23(nums []int, target int) []int {
	// binary search
	start := searchHelper(nums, target, true)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	/* v2_nohelper
	end := 0
	for i := start; i < len(nums) && nums[i] == target ; i++{
		end = i
	}
	return []int{start,end}
	*/
	return []int{start, searchHelper(nums, target, false) - 1}
	// v2_nohelper: Runtime: 8 ms, faster than 83.02% of Go online submissions for
	//		Find First and Last Position of Element in Sorted Array.
	// v3_helper: Runtime: 4 ms, faster than 98.94% of Go online submissions for
	//		Find First and Last Position of Element in Sorted Array.
	// v2_nohelper: Memory Usage: 4.1 MB, less than 50.00% of Go online submissions for
	//		Find First and Last Position of Element in Sorted Array.
}

func searchRange(nums []int, target int) []int {
	// v4: copied from best one
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if nums[0] > target {
		return []int{-1, -1}
	}

	if nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	low := 0
	high := len(nums) - 1
	for low <= high {
		median := (low + high) / 2

		if nums[median] == target {
			return getResult(nums, target, median)
		} else if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return []int{-1, -1}
	}

	if nums[low] == target {
		return getResult(nums, target, low)
	}

	return []int{-1, -1}
}

func getResult(nums []int, target, found int) []int {
	min := found
	max := found
	i := found + 1
	d := found - 1
	for i != -1 || d != -1 {
		if i != -1 {
			if i < len(nums) {
				if nums[i] == target {
					if max < i {
						max = i
					}
					i++
				} else {
					i = -1
				}
			} else {
				i = -1
			}
		}

		if d != -1 {
			if d >= 0 {
				if nums[d] == target {
					if min > d {
						min = d
					}
					d--
				} else {
					d = -1
				}
			} else {
				d = -1
			}
		}
	}

	return []int{min, max}
}
