// author: Gadigeppa
// date: 22-May-2020

// 2. 0-1 Knapsack
//
// Question: Given a list of items with values and weights, as well as a max weight,
// find the maximum value you can generate from items where the sum of the weights
// is less than the max.

// items  = {( w:1 ,  v:6 ), ( w:2 ,  v:10 ), ( w:3 ,  v:12 )}
// maxWeight =  5
// knapsack( items , maxWeight) =  22

// Time Complexity: O(2^n)
// Space Complexity: O(n)

// Algorithm Type: Recursion without memoization

package main

import "fmt"

func main() {

	items := [][]int{{1, 6},
		{2, 10},
		{3, 12}}
	maxWeight := 5

	fmt.Println(knapsack(items, maxWeight, 0, 0, 0))
}

func knapsack(items [][]int, maxWeight, index, currWeight, currValue int) int {

	if index == len(items) {
		return currValue
	}

	weight := items[index][0]
	value := items[index][1]

	with := 0

	if currWeight+weight <= maxWeight {
		with = knapsack(items, maxWeight, index+1, currWeight+weight, currValue+value)
	}

	without := knapsack(items, maxWeight, index+1, currWeight, currValue)

	return Max(with, without)

}

func Max(x, y int) int {

	if x > y {
		return x
	}
	return y
}
