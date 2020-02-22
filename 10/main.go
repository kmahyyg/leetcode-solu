package main

import (
	"fmt"
	"regexp"
)

func main() {
	
}

func isMatch(s string, p string) bool {
	d, _ := regexp.Match(fmt.Sprintf("^%s$", p), []byte(s))
	return d
}