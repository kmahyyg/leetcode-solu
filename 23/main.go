package main

import (
	"fmt"
	"sort"
)

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [[1,4,5],[1,3,4],[2,6]]
	lists1 := []*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6}},
	}
	dt := mergeKLists(lists1)
	for ; dt != nil; dt = dt.Next {
		fmt.Println(dt.Val)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res []int
	for _, v := range lists {
		for hd := v; hd != nil; hd = hd.Next {
			res = append(res, hd.Val)
		}
	}
	sort.Ints(res)
	newHd := &ListNode{}
	newStart := newHd
	for _, v := range res {
		newNd := &ListNode{Val: v}
		newHd.Next = newNd
		newHd = newHd.Next
	}
	return newStart.Next
}
