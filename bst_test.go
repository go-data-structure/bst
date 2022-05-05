package bst

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := New[int, string]()

	tree.InorderTraversal(ASC, func(k int, v string) {
		fmt.Println(k, ":", v)
	})

	_ = tree.Insert(3, "3")
	_ = tree.Insert(1, "1")
	_ = tree.Insert(8, "8")
	_ = tree.Insert(5, "5")
	_ = tree.Insert(6, "6")

	tree.InorderTraversal(ASC, func(k int, v string) {
		fmt.Println(k, ":", v)
	})
	tree.InorderTraversal(DESC, func(k int, v string) {
		fmt.Println(k, ":", v)
	})
}

func TestValid(t *testing.T) {
	var tree BST[int, string]
	_ = tree.Insert(3, "3")
	_ = tree.Insert(1, "1")
	_ = tree.Insert(8, "8")
	_ = tree.Insert(5, "5")
	_ = tree.Insert(6, "6")
	t.Log(tree.Get(7))
}