package bst

import "golang.org/x/exp/constraints"

func NewBst[K constraints.Ordered, V any]() *Node[K, V] {
	return &Node[K, V]{}
}
