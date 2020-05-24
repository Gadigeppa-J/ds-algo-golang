// author: Gadigeppa
// date: 24-May-2020

// 5. Consecutive Array
//
// Question:  Given an unsorted array, find the length of the longest sequence
// of consecutive numbers in the array.

// consecutive([ 4 ,  2 ,  1 ,  6 ,  5 ]) =  3 , [ 4 ,  5 ,  6 ]
// consecutive([ 5 ,  5 ,  3 ,  1 ]) =  1 , [ 1 ], [ 3 ], or [ 5 ]

// Time Complexity: O(n)
// Space Complexity: O(n)

// LeetCode: 128. Longest Consecutive Sequence

package main

import "fmt"

func main() {

	arr := []int{4, 2, 1, 6, 5}
	fmt.Println(consecutive(arr))

}

func consecutive(arr []int) int {

	set := make(map[int]bool)

	for _, val := range arr {
		set[val] = true
	}

	max, len := 0, 0

	for key, _ := range set {

		curr := key

		if _, ok := set[curr-1]; ok {
			continue
		}

		for ok := set[curr]; ok; ok = set[curr] {
			len++
			curr++
		}

		if len > max {
			max = len
		}

		len = 0

	}

	return max

}
