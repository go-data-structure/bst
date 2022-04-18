package bst

import "testing"

func TestInorderTraversal(t *testing.T) {
	tree := NewBst[int, string]()
	InorderTraversal(tree)

	tree, _ = Insert[int, string](tree, 3, "xxxx")
	tree, _ = Insert[int, string](tree, 1, "xxxx")
	tree, _ = Insert[int, string](tree, 8, "xxxx")
	tree, _ = Insert[int, string](tree, 5, "xxxx")
	tree, _ = Insert[int, string](tree, 6, "xxxx")

	InorderTraversal(tree)
	InorderTraversalDesc(tree)
}
