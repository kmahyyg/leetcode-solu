package main

func main() {
	
}

func getMinimumBuck(a int, b int) int{
	if a < b {return a}
	return b
}

func maxArea(height []int) int {
	maxdata := 0
	start := 0
	end := len(height) - 1
	for start < end {
		curdata := getMinimumBuck(height[start], height[end]) * (end - start)
		if curdata > maxdata { maxdata = curdata }
		if height[end] > height[start] {
			start++
		} else {
			end --
		}
	}
	return maxdata
}
