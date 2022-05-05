package bst

// LevelOrderTraversal level order traversal
func (n *Node[K, V]) LevelOrderTraversal(f func(key K, value V)) {
	var stacks []*Node[K, V]
	stacks = append(stacks, n)

	for len(stacks) > 0 {
		l := len(stacks)
		for i := 0; i < l; i++ {
			f(*stacks[i].Key, *stacks[i].Val)

			if stacks[i].Left != nil {
				stacks = append(stacks, stacks[i].Left)
			}
			if stacks[i].Right != nil {
				stacks = append(stacks, stacks[i].Right)
			}
		}
		stacks = stacks[l:]
	}
}
