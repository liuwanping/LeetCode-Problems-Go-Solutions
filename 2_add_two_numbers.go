package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := res
	flag := 0
	for l1 != nil || l2 != nil {
		sum := flag
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		flag = sum / 10
		q := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		p.Next = q
		p = q
	}
	if flag != 0{
		p.Next = &ListNode{
			Val: flag,
			Next: nil,
		}
	}
	return res.Next
}
