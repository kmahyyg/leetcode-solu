package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStrv1("hello", "ll"))
}

func strStrv1(haystack string, needle string) int {
	// official builtin library, rabin-karp alg. internally with bytes
	return strings.Index(haystack, needle)
}
