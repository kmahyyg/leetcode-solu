package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var i, j int
	for i, j = 0, 1; j < len(nums); {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}
