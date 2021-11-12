/*
	Definition for singly-linked list.
	type ListNode struct {
		Val int
		Next *ListNode
	}
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: -1}
	l3, carry := dummyNode, 0

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		carry = a + b + carry
		l3.Next = &ListNode{
			Val: carry % 10,
		}
		l3 = l3.Next
		carry = carry / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	// 判断最后一位是否有余位
	if carry > 0 {
		l3.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyNode.Next
}

func main() {
	fmt.Println("test code")
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)

	return
}
