// author: Gadigeppa
// date: 17-May-2020

// 20. Reverse Stack
//
// Question:  Given a stack, reverse the items without creating any additional data structures.
// reverse( 1 -> 2 -> 3 ) =  3 -> 2 -> 1

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// Algorithm Type: Recursion

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	s := stack{}

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)

	fmt.Println(s)

	reverse(&s)

	fmt.Println(s)

}

func reverse(s *stack) {

	if s.empty() {
		return
	}

	v, err := s.pop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reverse(s)
	pushToBottom(s, v)

}

func pushToBottom(s *stack, num int) {

	if s.empty() {
		s.push(num)
		return
	}

	v, err := s.pop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pushToBottom(s, num)
	s.push(v)

}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() (int, error) {

	if s.empty() {
		return -1, errors.New("Stack is empty!")
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return r, nil

}

func (s *stack) empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
