package main

import (
	"fmt"
)

func main() {
	testcases := []([]string) {
		[]string{"flower", "flow", "flight",},
		[]string{"dog", "racecar", "car",},
		[]string{"a", "a"},
		[]string{"c"},
	}

	for _,v := range testcases {
		fmt.Println(longestCommonPrefix(v))
	}
}

func longestCommonPrefix(strs []string) string {
	var restr = []uint8{}
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	} else {
		var minLen = len(strs[0])
		for _, v1 := range strs{
			curLen := len(v1)
			if curLen < minLen { minLen = curLen }
		}
		for i := 0; i < minLen; i++ {
			for j := 1; j < len(strs); j++ {
				 if strs[0][i] != strs[j][i] {
					return string(restr)
				 }
			}
			restr = append(restr, strs[0][i])
		}
		return string(restr)
	}
}