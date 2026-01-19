package main


import "fmt"

// Node represents the components of binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int){
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take a key value
// and Return true if there is a node in that value
func (n *Node) Search(k int) bool {
	
	if n.Key == k {
		return true
	}

	if n.Key < k {
		// move right
		if n.Right == nil {
			return false
		}
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			return false
		}
		return n.Left.Search(k)
	}

	return false
}


func main(){
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(10)
	tree.Insert(110)
	fmt.Println(tree)
	fmt.Println(tree.Search(20))
	fmt.Println(tree.Search(100))
}