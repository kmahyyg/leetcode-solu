package main

import (
	"fmt"
)

func main() {
	testcases := []string{
		"aaabaaaa",
		"abcdcba1234567890-=jtkpqpktjlzhs",
		"a",
		"bbbbbbb",
		"ab",
		"cc",
		"acbd",
		"cbbd",
		"aaaa",
	}
	for _, v := range testcases {
		fmt.Println(longestPalindrome(v))
	}
}

func expandAroundCenter(dt []byte, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(dt) && dt[L] == dt[R] {
		L--
		R++
	}
	return R - L - 1
}

func longestPalindrome(s string) string {
	// expand from center approach
	if len(s) <= 1 {
		return s
	}
	// byte array always has better performance
	dt := []byte(s)
	var lenf, start, end int
	for i := 0; i < len(dt); i++ {
		// take each char in original data as center char
		// consider the original data length is odd or even.
		// if length is odd
		len1 := expandAroundCenter(dt, i, i)
		// if length is even
		len2 := expandAroundCenter(dt, i, i+1)
		// get the maximum palindrome length
		if len1 > len2 {
			lenf = len1
		} else {
			lenf = len2
		}
		// if length large than pre-gotten one, change it.
		if lenf > end-start {
			// length = start to end
			// center char is i
			start = i - (lenf-1)/2
			end = i + lenf/2
		}
	}
	// convert back and go ahead
	return string(dt[start : end+1])
}
