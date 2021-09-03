package bst

import "testing"

func TestValid(t *testing.T) {
	bst := &Node{Val: 2}
	bst.Left = &Node{Val: 1}
	bst.Right = &Node{Val: 3}

	if Valid(bst) {
		t.Log("true")
	}
}
