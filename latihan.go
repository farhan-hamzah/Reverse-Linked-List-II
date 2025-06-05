package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Fungsi utama untuk membalik sebagian linked list
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// 1. Geser prev ke node sebelum 'left'
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// 2. Mulai membalik dari prev.Next (node ke-left)
	current := prev.Next
	var next *ListNode

	// 3. Balik pointer antara left dan right
	for i := 0; i < right-left; i++ {
		next = current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// Fungsi bantu untuk membuat linked list dari slice
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Fungsi bantu untuk mencetak linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// Fungsi main untuk menjalankan contoh
func main() {
	values := []int{1, 2, 3, 4, 5}
	head := createLinkedList(values)
	fmt.Print("Original list: ")
	printLinkedList(head)

	left, right := 2, 4
	newHead := reverseBetween(head, left, right)

	fmt.Printf("List after reversing from %d to %d: ", left, right)
	printLinkedList(newHead)
}
