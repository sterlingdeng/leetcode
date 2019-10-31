package data_structures

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}
