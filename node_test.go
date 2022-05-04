package bst

import (
	"testing"
)

func TestValid(t *testing.T) {
	var tree BST[int, string]
	tree.Insert(1, "xxxx")
	t.Log(tree.Get(1))
}
