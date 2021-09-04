package bst

// BST is binary search tree.
type BST struct {
	Root *Node
}

// Node is binary search tree node.
type Node struct {
	Left  *Node
	Right *Node
	ID    int
	Val   interface{}
}
