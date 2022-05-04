package bst

import "golang.org/x/exp/constraints"

func NewBst[K constraints.Ordered, V any]() *Node[K, V] {
	var n *Node[K, V]
	return n
}

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
func Insert[K constraints.Ordered, V any](t *Node[K, V], key K, value V) (*Node[K, V], error) {
	if t == nil {
		node := newBSTNode[K, V](key, value)
		return &node, nil
	}

	var cur = t

	for {
		if *cur.Key == key {
			return nil, ErrKeyExists
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

	return t, nil
}
