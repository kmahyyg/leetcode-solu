package main

/**
 * Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var headNode = &ListNode{
		Val:  0,
		Next: nil,
	}
	var nowNode = headNode
	var curr = 0
	var cFlag = 0
	var eFlag1, eFlag2 bool
	for true {
		eFlag1 = l1 == nil
		eFlag2 = l2 == nil
		if eFlag1 && eFlag2 {
			if cFlag != 0 {
				nowNode.Next = &ListNode{Val: cFlag}
			}
			return headNode.Next
		}
		if eFlag1 {
			l1 = &ListNode{Val: 0}
		}
		if eFlag2 {
			l2 = &ListNode{Val: 0}
		}
		suma := l1.Val + l2.Val
		if cFlag != 0 {
			suma += cFlag
		}
		curr = suma % 10
		cFlag = suma / 10
		l1 = l1.Next
		l2 = l2.Next
		nowNode.Next = &ListNode{Val: curr}
		nowNode = nowNode.Next
	}
	return headNode.Next
}
