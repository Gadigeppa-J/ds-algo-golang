// author: Gadigeppa
// date: 15-May-2020

// 23. Print Reversed Linked List
//
// Question: Given a linked list, write a function that prints the nodes of the list in reverse order.
// printReversedList(1 -> 2 -> 3)
// 3
// 2
// 1

// Time Complexity: O(n)
// Space Complexity: O(n)
// Using recursion

package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {

	n1 := node{val: 1}
	n2 := node{val: 2}
	n3 := node{val: 3}
	n4 := node{val: 4}
	n5 := node{val: 5}

	n1.next = &n2
	n2.next = &n3
	n3.next = &n4
	n4.next = &n5

	printReversedList(&n1)

}

func printReversedList(h *node) {
	if h == nil {
		return
	}

	printReversedList(h.next)
	fmt.Println(h.val)
}
