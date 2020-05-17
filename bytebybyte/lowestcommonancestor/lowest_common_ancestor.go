// author: Gadigeppa
// date: 17-May-2020

// 18. Lowest Common Ancestor
//
// Question:  Given two nodes in a binary tree, write a function to find the lowest common ancestor.
//
//         1
//       /   \
//     3       4
//   /  \     / \
// 4     5  6    7
//
// lcs( 4 ,  3 ) =  1
// lcs( 6 ,  7 ) =  3

// Time Complexity: O(n)
// Space Complexity: O(h) where h is the height of the binary tree

// Algorithm Type: Recursion

package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func main() {

	n1 := node{val: 1}
	n2 := node{val: 2}
	n3 := node{val: 3}
	n4 := node{val: 4}
	n5 := node{val: 5}
	n6 := node{val: 6}
	n7 := node{val: 7}

	n1.left = &n2
	n1.right = &n3

	n2.left = &n4
	n2.right = &n5

	n3.left = &n6
	n3.right = &n7

	fmt.Println(lcs(&n1, 6, 7))
}

func lcs(h *node, x, y int) int {

	if h == nil {
		return -1
	}

	if (*h).val == x || (*h).val == y {
		return (*h).val
	}

	left := lcs((*h).left, x, y)
	right := lcs((*h).right, x, y)

	if left != -1 && right != -1 {
		return (*h).val
	} else if left != -1 {
		return left
	} else if right != -1 {
		return right
	} else {
		return -1
	}

}
