package main

import "fmt"

func main() {
	testcases := []int{3, 4, 5}
	for _, v1 := range testcases {
		dt := generateParenthesis(v1)
		for k2, v2 := range dt {
			fmt.Printf("%d : %s , ", k2, v2)
		}
		fmt.Printf("\n")
	}
}

func generateParenthesisHelper(data string, opens int, closes int, n int, result *[]string) {
	if len(data) == n*2 {
		*result = append(*result, data)
		return
	}
	if opens < n {
		generateParenthesisHelper(data+"(", opens+1, closes, n, result)
	}
	if closes < opens {
		generateParenthesisHelper(data+")", opens, closes+1, n, result)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisHelper("", 0, 0, n, &res)
	return res
}
