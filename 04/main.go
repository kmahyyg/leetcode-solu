package main

import "sort"

func main(){

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums3 := append(nums1, nums2...)
    sort.Ints(nums3)
    dtlen := len(nums3)
    if dtlen % 2 == 0 {
        return float64(nums3[dtlen / 2] + nums3[dtlen / 2 - 1]) / 2
    } else {
        return float64(nums3[dtlen / 2])
    }
}

