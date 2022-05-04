package bst

import "golang.org/x/exp/constraints"

// Node is binary search tree node.
type Node[K constraints.Ordered, V any] struct {
	Left  *Node[K, V]
	Right *Node[K, V]
	Key   *K
	Val   *V
}

func newBSTNode[K constraints.Ordered, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key: &key,
		Val: &value,
	}
}

// Get value form bst
func (n *Node[K, V]) Get(key K) V {
	var cur = n
	for {
		if *cur.Key == key {
			return *cur.Val
		} else if *cur.Key > key {
			cur = cur.Right
		} else if *cur.Key < key {
			cur = cur.Left
		}
	}
}

// Insert insert a key-value to bst
func (n *Node[K, V]) Insert(key K, value V) error {
	var cur = n

	for {
		if *cur.Key == key {
			return ErrKeyExists
		}

		if *cur.Key > key {
			if cur.Left == nil {
				cur.Left = newBSTNode[K, V](key, value)
				break
			}
			cur = cur.Left
		} else if *cur.Key < key {
			if cur.Right == nil {
				cur.Right = newBSTNode[K, V](key, value)
				break
			}
			cur = cur.Right
		}
	}

	return nil
}

// Upsert a key-value to bst
func (n *Node[K, V]) Upsert(key K, value V) {
	var cur = n

	for {
		if *cur.Key == key {
			*cur.Val = value
			return
		}

		if *cur.Key > key {
			if cur.Left == nil {
				cur.Left = newBSTNode[K, V](key, value)
				break
			}
			cur = cur.Left
		} else if *cur.Key < key {
			if cur.Right == nil {
				cur.Right = newBSTNode[K, V](key, value)
				break
			}
			cur = cur.Right
		}
	}
}

// Delete node from bst
func Delete[K constraints.Ordered, V any](t *Node[K, V], key K) error {
	return nil
}
