package main

import "fmt"

func main() {
	fmt.Println(divide(242342, 12))
}

const (
	MaxInt32     = 1<<31 - 1
	MinInt32     = -1 << 31
	HalfMinInt32 = -1073741824
)

// The methodlogy of programming is: do not rely on the unspecific behavior of compiler and system,
// and try to avoid overflow as much as possible.
func divide(dividend int, divisor int) int {
	if dividend == MinInt32 && divisor == -1 {
		return MaxInt32
	}
	if dividend == MinInt32 && divisor == 1 {
		return MinInt32
	}

	// negative sign
	negatives := 2
	// transform to negative to have larger range to avoid overflow
	if dividend > 0 {
		negatives -= 1
		dividend = -dividend
	}
	if divisor > 0 {
		negatives -= 1
		divisor = -divisor
	}

	// find the largest doubling of divisor
	// if less than HalfMinInt32, it would overflowed
	// then we just stop
	maxBit := 0
	for divisor >= HalfMinInt32 && divisor+divisor >= dividend {
		maxBit += 1
		divisor += divisor
	}

	quotient := 0
	for bit := maxBit; bit >= 0; bit-- {
		// just as normal long division in our life
		if divisor >= dividend {
			quotient -= (1 << bit)
			dividend -= divisor
		}
		// just like two's complement, something like indirect abs()
		divisor = (divisor + 1) >> 1
	}

	// transform back to positive
	if negatives != 1 {
		quotient = -quotient
	}
	return quotient
}
