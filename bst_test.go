package bst

import (
	"testing"
)

func TestValid(t *testing.T) {
	tree := NewBst[int, string]()
	tree, _ = Insert[int, string](tree, 1, "xxxx")
	t.Log(tree.Get(1))
}
