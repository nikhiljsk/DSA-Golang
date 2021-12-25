package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func print(head *ListNode) {
	iter := head
	for iter != nil {
		fmt.Print(iter.Val, "->")
		iter = iter.Next
	}
	fmt.Println()
}

func duplicate(head *ListNode) *ListNode {
	iter := head
	iter = iter.Next
	newHead := &ListNode{Next: nil, Val: head.Val}
	newTail := newHead
	for iter != nil {
		newNode := &ListNode{}
		newNode.Val = iter.Val
		newTail.Next = newNode
		newTail = newNode
		newTail.Next = nil
		iter = iter.Next
	}
	return newHead
}

func main() {
	n5 := &ListNode{Next: nil, Val: 5}
	n4 := &ListNode{Next: n5, Val: 4}
	n3 := &ListNode{Next: n4, Val: 3}
	n2 := &ListNode{Next: n3, Val: 2}
	n1 := &ListNode{Next: n2, Val: 1}
	fmt.Println("First LL:")
	print(n1)
	newList := duplicate(n1)
	fmt.Println("Cloned LL:")
	print(newList)
}
