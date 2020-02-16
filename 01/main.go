package main

func twoSum_perfect(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := index[target-num]; ok == true {
			return []int{j, i}
		}
		index[num] = i
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
