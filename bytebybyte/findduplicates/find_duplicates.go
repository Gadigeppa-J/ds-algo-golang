// author: Gadigeppa
// date: 22-May-2020

// 4. Find Duplicates
//
// Question:  Given an array of integers where each value 1 <= x <= len(array),
// write a function that finds all the duplicates in the array.

// dups([ 1 ,  2 ,  3 ]) = []
// dups([ 1 ,  2 ,  2 ]) = [2]
// dups([ 3 ,  3 ,  3 ]) = [3]
// dups([ 2 ,  1 ,  2 ,  1 ]) = [ 1 ,  2 ]

// LeetCode: 442. Find All Duplicates in an Array

// Time Complexity: O(n)
// Space Complexity: O(1)

package main

import (
	"fmt"
)

func main() {

	// using array and not slice of int
	// so that the original array is not modified
	a := [4]int{2, 1, 2, 1}
	fmt.Println(dups(a))

}

func dups(a [4]int) []int {

	set := make(map[int]bool)
	result := []int{}

	for _, v := range a {

		di := Abs(v) - 1

		if a[di] < 0 {
			set[Abs(v)] = true
		} else {
			a[di] = -a[di]
		}
	}

	for k := range set {
		result = append(result, k)
	}

	return result

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
