package main

import "fmt"

func main() {
	testcases := []string{
		"MCMXCIV",
		"III",
		"IV",
		"IX",
		"LVIII",
	}
	for _,v := range testcases{
		fmt.Println(romanToInt(v))
	}
}

func romanToInt(s string) int {
	st := []byte(s)
	mapTable := map[byte]int{73: 1, 86: 5, 88:10, 76:50, 67:100, 68: 500, 77: 1000}
	curPos := len(st) - 1
	res := 0
	for curPos >= 0{
		if curPos - 1 >= 0 && mapTable[st[curPos]] > mapTable[st[curPos - 1]]{
			res = res + mapTable[st[curPos]] - mapTable[st[curPos-1]]
			curPos -= 2
			continue
 		} else {
 			res += mapTable[st[curPos]]
 			curPos -= 1
 			continue
		}
	}
	return res
}