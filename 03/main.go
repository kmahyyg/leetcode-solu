package main

import "fmt"

func main() {
	fmt.Printf("%d", lengthOfLongestSubstring("abcjk1234abcjti12345"))
}

func lengthOfLongestSubstring(s string) int {
	// transfer to byte array, improve performance
	oriData := []byte(s)
	// ascii Table as a sliding window recorder
	asciiTable := make([]int, 128)
	// original data length
	dtLen := len(oriData)
	// temporary variable
	var maxL, windLen int
	// start for iterate the whole data bytearray
	for startPt, j := 0, 0; startPt < dtLen && j < dtLen; j++ {
		// if corresponding location counter value larger than start point
		// which means there's a repeat char
		if asciiTable[oriData[j]] > startPt {
			// change the sliding window start point to next char
			// contrast window
			startPt = asciiTable[oriData[j]]
		}
		// windows length = current index - start point + 1
		windLen = j - startPt + 1
		// exchange max value
		if windLen > maxL {
			maxL = windLen
		}
		// record data location in sliding windows table, expand wind
		asciiTable[oriData[j]] = j + 1
	}
	return maxL
}
