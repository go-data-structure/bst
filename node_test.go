package bst

import (
	"testing"
)

func TestValid(t *testing.T) {
	tree := NewBst[int, string]()
	_ = tree.Insert(1, "xxxx")
	t.Log(tree.Get(1))
}
