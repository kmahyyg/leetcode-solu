package main

import (
	"fmt"
	"math"
)

func main() {
	testcases := []string{
		"21",
		"     32",
		"    words and 32",
		"-382472347283472",
		"324 with fucking",
		"+afu234pl",
		"+-2",
		"  +0   123",
		" +8 123",
		"9223372036854775808",
		"  -42",
		"-91283472332",
		"-    234",
		"0-1",
	}
	for _,v := range testcases {
		fmt.Println(myAtoi(v))
	}
}

func myAtoi(str string) int {
	var res int
	finn := false
	marked := false
	sign := 1
	tbarray := []byte(str)
	for _,v := range tbarray {
		if v == 32 && !finn && !marked {
			continue
		} else if v <= 57 && v >= 48{
			if res * sign > math.MaxInt32 {
				return math.MaxInt32
			} else if res * sign < math.MinInt32 {
				return math.MinInt32
			}
			if finn {
				res = res * 10 + int(v) - 48
			} else {
				res = int(v) - 48
				finn = true
			}
		} else if v == 43 && !marked && !finn{
			marked = true
			sign = 1
		} else if v == 45 && !marked && !finn{
			marked = true
			sign = -1
		} else {
			break
		}
	}
	if res * sign > math.MaxInt32 {
		return math.MaxInt32
	} else if res * sign < math.MinInt32 {
		return math.MinInt32
	}
	return res * sign
}