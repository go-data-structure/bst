package bst

import "golang.org/x/exp/constraints"

// BST binary search tree
type BST[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

// New a binary search tree
func New[K constraints.Ordered, V any]() *BST[K, V] {
	return &BST[K, V]{}
}

// Get value form bst
func (t *BST[K, V]) Get(key K) V {
	if t.root == nil {
		var null V
		return null
	}

	v, _ := t.root.Get(key)
	return v
}

// Insert insert a key-value to bst
func (t *BST[K, V]) Insert(key K, value V) error {
	if t.root == nil {
		t.root = newBSTNode[K, V](key, value)
		return nil
	}
	return t.root.Insert(key, value)
}

// Upsert a key-value to bst
func (t *BST[K, V]) Upsert(key K, value V) {
	if t.root == nil {
		t.root = newBSTNode[K, V](key, value)
		return
	}

	t.root.Upsert(key, value)
}

// InorderTraversal inorder traversal
func (t *BST[K, V]) InorderTraversal(orderDirect OrderDirect, f func(key K, value V)) {
	if t.root == nil {
		return
	}

	switch orderDirect {
	case ASC:
		t.root.InorderTraversalAsc(f)
	case DESC:
		t.root.InorderTraversalDesc(f)
	}
}
