package bst

// Valid valid binary search tree.
func Valid(root *Node) bool {
	var validBST func(n, min, max *Node) bool
	validBST = func(n, min, max *Node) bool {
		if n == nil {
			return true
		}

		if min != nil && n.Val <= min.Val {
			return false
		}
		if max != nil && n.Val >= max.Val {
			return false
		}

		return validBST(n.Left, min, n) && validBST(n.Right, n, max)
	}
	return validBST(root, nil, nil)
}
