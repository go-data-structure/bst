package bst

// InorderTraversalAsc inorder traversal in asc
func (n *Node[K, V]) InorderTraversalAsc(f func(key K, value V)) {
	stacks := make([]*Node[K, V], 0)
	var cur = n
	var hasNext = true

	for hasNext {
		for n := cur; n != nil; n = n.Left {
			stacks = append(stacks, n)
		}

		cur = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		f(*cur.Key, *cur.Val)

		cur = cur.Right

		if cur == nil && len(stacks) == 0 {
			hasNext = false
		}
	}
}

// InorderTraversalDesc inorder traversal in desc
func (n *Node[K, V]) InorderTraversalDesc(f func(key K, value V)) {
	stacks := make([]*Node[K, V], 0)
	var cur = n
	var hasNext = true

	for hasNext {
		for n := cur; n != nil; n = n.Right {
			stacks = append(stacks, n)
		}

		cur = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		f(*cur.Key, *cur.Val)

		cur = cur.Left

		if cur == nil && len(stacks) == 0 {
			hasNext = false
		}
	}
}
