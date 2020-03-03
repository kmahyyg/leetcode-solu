package main

import (
	"fmt"
	"sort"
)

func main() {
	l1test := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2test := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	retVal := mergeTwoListsVersion2(l1test, l2test)
	for retVal.Next != nil {
		fmt.Printf("%#v \n", retVal)
		retVal = retVal.Next
	}

}

// ListNode :Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Version 1
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.8 MB, less than 7.14% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoListsVersion1(l1 *ListNode, l2 *ListNode) *ListNode {
	var dt []int
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l2 == nil && l1 != nil {
		return l1
	} else {
		for hd := l1; hd != nil; hd = hd.Next {
			dt = append(dt, hd.Val)
		}
		for hd := l2; hd != nil; hd = hd.Next {
			dt = append(dt, hd.Val)
		}
		sort.Ints(dt)
		newLst := &ListNode{}
		newLstStart := newLst
		for _, v := range dt {
			newLst.Next = &ListNode{Val: v, Next: nil}
			newLst = newLst.Next
		}
		return newLstStart.Next
	}
}

// Version 2
// HelperFunction to Select Smaller Ones

func selectSmaller(lnode1 *ListNode, lnode2 *ListNode) (*ListNode, int) {
	if lnode1 == nil && lnode2 != nil {
		return lnode2, 2
	}
	if lnode1 != nil && lnode2 == nil {
		return lnode1, 1
	}
	if lnode1 == nil && lnode2 == nil {
		return nil, -1
	}
	if lnode1.Val < lnode2.Val {
		return lnode1, 1
	}
	return lnode2, 2
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHd := &ListNode{}
	newHdStart := newHd
	for l1 != nil || l2 != nil {
		var casedata int
		newHd.Next, casedata = selectSmaller(l1, l2)
		newHd = newHd.Next
		switch casedata {
		case 1:
			l1 = l1.Next
		case 2:
			l2 = l2.Next
		}
	}
	return newHdStart.Next
}
