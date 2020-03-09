package main

func main() {

}

const (
	MaxInt32     = 1<<31 - 1
	MinInt32     = -1 << 31
	HalfMinInt32 = -1073741824
	HalfMaxInt32 = 1073741824
)

// The methodlogy of programming is: do not rely on the unspecific behavior of compiler and system,
// and try to avoid overflow as much as possible.
func divide(dividend int, divisor int) int {
	
}
