package main

import (
	"fmt"
	"sort"
)

func main(){
	datas := []([]int){
		[]int{3,6,2,3},
		[]int{3,2,3,4},
		[]int{1,2,1},
		[]int{2,1,2},
	}
	for _,dt := range datas{
		fmt.Println(largestPerimeter(dt))
	}
}

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A)-1; i >= 0; i-- {
		if i-2 < 0{
			return 0
		}
		if A[i-1] + A[i-2] > A[i] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

