package main

import "fmt"

func main() {
	testcases := []string{
		"",
		"1234",
		"44552",
		"231",
	}
	for _, v := range testcases {
		fmt.Println(letterCombinations(v))
	}
}

func letterCombinationsHelper(digits string, index int, combs []byte, res *[]string) {
	// from 0 to 9, 0 and 1 is nothing
	lookupTb := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if index == len(digits) {
		*res = append(*res, string(combs))
		return
	}
	curWord := lookupTb[digits[index]-'0']
	for i := 0; i < len(curWord); i++ {
		letterCombinationsHelper(digits, index+1, append(combs, curWord[i]), res)
	}
}

func letterCombinations(digits string) (res []string) {
	if len(digits) == 0 {
		return
	}
	letterCombinationsHelper(digits, 0, []byte{}, &res)
	return
}
