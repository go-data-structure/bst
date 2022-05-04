package bst

import "golang.org/x/exp/constraints"

// Node is binary search tree node.
type Node[K constraints.Ordered, V any] struct {
	Left  *Node[K, V]
	Right *Node[K, V]
	Key   *K
	Val   *V
}

func newBSTNode[K constraints.Ordered, V any](key K, value V) Node[K, V] {
	return Node[K, V]{
		Key: &key,
		Val: &value,
	}
}

// Get value form bst
func (t *Node[K, V]) Get(key K) V {
	var cur = t
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
func (t *Node[K, V]) Insert(key K, value V) error {
	if t == nil {
		*t = newBSTNode[K, V](key, value)
		return nil
	}

	var cur = t

	for {
		if *cur.Key == key {
			return ErrKeyExists
		}

		if *cur.Key > key {
			if cur.Left == nil {
				node := newBSTNode[K, V](key, value)
				cur.Left = &node
				break
			}
			cur = cur.Left
		} else if *cur.Key < key {
			if cur.Right == nil {
				node := newBSTNode[K, V](key, value)
				cur.Right = &node
				break
			}
			cur = cur.Right
		}
	}

	return nil
}

// Delete node from bst
func Delete[K constraints.Ordered, V any](t *Node[K, V], key K) error {
	return nil
}
