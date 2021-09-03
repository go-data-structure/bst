package bst

// Node is binary search tree node.
type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

// BST is binary search tree.
type BST struct {
	Root *Node
}
