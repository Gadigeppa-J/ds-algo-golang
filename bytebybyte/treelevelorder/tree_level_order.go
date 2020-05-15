
// author: Gadigeppa
// date: 15-May-2020

// 44. Tree Level Order
//
// Question: Given a tree, write a function that prints out the nodes of the tree in level order
// Time Complexity: O(n)
// Space Complexity: O(n)

package main

import "fmt"

type node struct{
  val int
  left *node
  right *node
}

func main(){

  n1 := node{val:1}
  n2 := node{val:2}
  n3 := node{val:3}
  n4 := node{val:4}
  n5 := node{val:5}
  n6 := node{val:6}
  n7 := node{val:7}

  n1.left = &n2
  n1.right = &n3

  n2.left = &n4
  n2.right = &n5

  n3.left = &n6
  n3.right = &n7

  traverse(&n1)

}

func traverse(r *node){

  q := []*node{r}

  for len(q) != 0 {

      c := q[0]
      q = q[1:]

      fmt.Println(c.val)

      if c.left != nil{
        q=append(q, c.left)
      }

      if c.right != nil{
        q=append(q, c.right)
      }

  }

}
