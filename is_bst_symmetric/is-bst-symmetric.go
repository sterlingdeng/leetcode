package is_bst_symmetric

import (
	. "hack_hour/data_structures"
)

func IsBstSymmetric(t *BinarySearchTree) bool {
	if t == nil {
		return false
	}
	return isSym(t, t)
}

func isSym(l, r *BinarySearchTree) bool {
	if l == nil && r == nil {
		return true
	}

	if (l == nil && r != nil) || (l != nil && r == nil)  {
		return false
	}

	if l.Value != r.Value {
		return false
	}

	return isSym(l.Left, r.Right) && isSym(l.Right, r.Left)
}