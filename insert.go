package bst

import "golang.org/x/exp/constraints"

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
