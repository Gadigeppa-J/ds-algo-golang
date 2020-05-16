// author: Gadigeppa
// date: 16-May-2020

// 3. Matrix product
//
// Question:  Given a matrix, find the path from top left to bottom
// right with the greatest product by moving only down and right.
// nxn matrix

// [ 1 ,  2 ,  3 ]
// [ 4 ,  5 ,  6 ]
// [ 7 ,  8 ,  9 ]
// 1  ->  4  ->  7  ->  8  ->  9 2016

// [ -1 ,  2 ,  3 ]
// [ 4 ,  5 ,  -6 ]
// [ 7 ,  8 ,  9 ]
// -1  ->  4  ->  5  ->  -6  ->  9 1080

// Time Complexity: O(2^n)
// Worst case it creates two branches per recursive call i.e down and right

// Space Complexity: O(n)
// At most 2n stacks exits in stack call at any given point of time

// Algorithm Type: Recursion without memoization

package main

import "fmt"

func main() {

	m := [][]int{
		{-1, 2, 3},
		{4, 5, -6},
		{7, 8, 9},
	}

	fmt.Println(product(m, 0, 0, 3, 1))

}

func product(m [][]int, i, j, n, p int) int {

	if (i == n-1) && (j == n-1) {
		return p * m[i][j]
	}

	p *= m[i][j]

	down := 0
	right := 0

	if i < n-1 {
		down = product(m, i+1, j, n, p)
	}

	if j < n-1 {
		right = product(m, i, j+1, n, p)
	}

	return Max(down, right)

}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
