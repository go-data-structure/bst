package bst

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// InorderTraversal inorder traversal in asc
func InorderTraversal[K constraints.Ordered, V any](t *Node[K, V]) {
	if t == nil {
		return
	}

	stacks := make([]*Node[K, V], 0)
	var cur = t
	var hasNext = true

	for hasNext {
		for n := cur; n != nil; n = n.Left {
			stacks = append(stacks, n)
		}

		cur = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		fmt.Println(*cur.Key)
		fmt.Println(*cur.Val)

		cur = cur.Right

		if cur == nil && len(stacks) == 0 {
			hasNext = false
		}
	}
}

// InorderTraversalDesc inorder traversal in desc
func InorderTraversalDesc[K constraints.Ordered, V any](t *Node[K, V]) {
	if t == nil {
		return
	}

	stacks := make([]*Node[K, V], 0)
	var cur = t
	var hasNext = true

	for hasNext {
		for n := cur; n != nil; n = n.Right {
			stacks = append(stacks, n)
		}

		cur = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		fmt.Println(*cur.Key)
		fmt.Println(*cur.Val)

		cur = cur.Left

		if cur == nil && len(stacks) == 0 {
			hasNext = false
		}
	}
}
