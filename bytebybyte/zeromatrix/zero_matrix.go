// author: Gadigeppa
// date: 24-May-2020

// 6. Zero Matrix
//
// Question:  Given a boolean matrix, update it so that if any cell is true,
// all the cells in that row and column are true.

// [true,   false ,   false ]         [true,   true ,   true ]
// [false,   false ,   false ]   ->   [true,   false ,   false ]
// [false,   false ,   false ]        [true,   false ,   false ]

// Time Complexity: O(mn)
// Space Complexity: O(1)

// LeetCode: 73. Set Matrix Zeroes

package main

import "fmt"

func main() {

	matrix := [][]bool{{true, false, false},
		{false, false, false},
		{false, false, false}}

	zeroMatrix(matrix)

	fmt.Println(matrix)

}

func zeroMatrix(matrix [][]bool) {

	rwTrue := false
	clTrue := false

	for j, _ := range matrix[0] {
		if matrix[0][j] {
			rwTrue = true
			break
		}
	}

	for i, _ := range matrix {
		if matrix[i][0] {
			clTrue = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {

		for j := 1; j < len(matrix[0]); j++ {

			if matrix[i][j] {
				matrix[i][0] = true
				matrix[0][j] = true
			}

		}

	}

	for j := 1; j < len(matrix[0]); j++ {

		if matrix[0][j] {

			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = true
			}

		}

	}

	for i := 1; i < len(matrix); i++ {

		if matrix[i][0] {

			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = true
			}

		}

	}

	if rwTrue {

		for j, _ := range matrix[0] {
			matrix[0][j] = true

		}

	}

	if clTrue {

		for i, _ := range matrix {
			matrix[i][0] = true

		}

	}

}
