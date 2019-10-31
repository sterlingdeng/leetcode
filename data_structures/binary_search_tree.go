package data_structures

type BinarySearchTree struct {
	Value int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func (b *BinarySearchTree) Add(n int) {
	add(n, b)
}

func add(n int, b *BinarySearchTree) *BinarySearchTree {
	if b == nil {
		return &BinarySearchTree{
			Value: n,
		}
	}
	if n < b.Value {
		b.Left = add(n, b.Left)
	} else {
		b.Right = add(n, b.Right)
	}
	return b
}

func (b *BinarySearchTree) Contains(n int) bool {
	return contains(n, b)
}

func contains(n int, b *BinarySearchTree) bool {
	if b == nil {
		return false
	}
	if n == b.Value {
		return true
	}
	if n < b.Value {
		return contains(n, b.Left)
	}
	return contains(n, b.Right)
}

func (b *BinarySearchTree) Height() int {
	return height(b)
}

func height(node *BinarySearchTree) int {
	if node == nil {
		return 0
	}
	leftHeight := 1 + height(node.Left)
	rightHeight := 1 + height(node.Right)
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}


