// author: Gadigeppa
// date: 16-May-2020

// 11. Zero Sum Subarray
//
// Question:  Given an array, write a function to find any subarray that sums to zero, if one exists.
// zeroSum({ 1 ,  2 ,  -5 ,  1 ,  2 ,  -1 }) = [ 2 ,  -5 ,  1 ,  2 ]

// Time Complexity: O(n)
// Space Complexity: O(n)

package main

import "fmt"

func main() {
	arr := []int{1, 2, -5, 1, 2, -1}
	fmt.Println(zeroSum(arr))
}

func zeroSum(arr []int) []int {

	m := make(map[int]int)
	sum := 0
	res := []int{}

	for i, v := range arr {
		sum += v
		if mv, ok := m[sum]; ok {
			for j := mv + 1; j <= i; j++ {
				res = append(res, arr[j])
			}
			return res
		} else if sum == 0 {
			for j := 0; j <= i; j++ {
				res = append(res, arr[j])
			}
			return res
		} else {
			m[sum] = i
		}
	}

	return res
}
