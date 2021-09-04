package bst

// Valid valid binary search tree.
func Valid(root *Node) bool {
	var validBST func(n, min, max *Node) bool
	validBST = func(n, min, max *Node) bool {
		if n == nil {
			return true
		}

		if min != nil && n.ID <= min.ID {
			return false
		}
		if max != nil && n.ID >= max.ID {
			return false
		}

		return validBST(n.Left, min, n) && validBST(n.Right, n, max)
	}
	return validBST(root, nil, nil)
}
