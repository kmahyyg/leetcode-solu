package main

import "fmt"

func main() {
	testcases := []int{
		123,
		-123,
		120,
		2147483649,
	}
	for _, v := range testcases {
		fmt.Println(reverse(v))
	}
}

func reverse(x int) int {
	var res int
	if x == 0 {
		return 0
	}
	for x != 0 {
		res = res*10 + x%10
		if res > (1<<31-1) || res < -1<<31 {
			return 0
		}
		x /= 10
	}
	return res
}
