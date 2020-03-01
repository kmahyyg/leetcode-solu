package main

import "fmt"

func main() {
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	for j := removeNthFromEnd(&node1, 2); j != nil; j = j.Next {
		fmt.Printf("%#v\n", j)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// n is always valid.
	nHead := &ListNode{0, nil}
	marker, targetPrev := nHead, nHead
	marker.Next = head
	for i := 0; i <= n; i++ {
		targetPrev = targetPrev.Next
	}

	for targetPrev != nil {
		marker = marker.Next
		targetPrev = targetPrev.Next
	}
	marker.Next = marker.Next.Next

	return nHead.Next
}
