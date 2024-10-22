package main

import (
	"errors"
	"fmt"
)

// Node represents a single node in the binary tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// FullBinaryTree represents a full binary tree.
type FullBinaryTree struct {
	Root *Node
}

// NewFullBinaryTree creates a new full binary tree.
func NewFullBinaryTree(root *Node) *FullBinaryTree {
	return &FullBinaryTree{Root: root}
}

// Insert inserts a new node into the tree.
func (fbt *FullBinaryTree) Insert(value int) error {
	if fbt.Root == nil {
		fbt.Root = &Node{Value: value}
		return nil
	}

	var queue []*Node
	queue = append(queue, fbt.Root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil {
			node.Left = &Node{Value: value}
			return nil
		}

		if node.Right == nil {
			node.Right = &Node{Value: value}
			return nil
		}

		queue = append(queue, node.Left, node.Right)
	}

	return errors.New("tree is full")
}

// PrintTree prints the tree using in-order traversal.
func (fbt *FullBinaryTree) PrintTree(node *Node) {
	if node != nil {
		fbt.PrintTree(node.Left)
		fmt.Print(node.Value, " ")
		fbt.PrintTree(node.Right)
	}
}

func main() {
	tree := NewFullBinaryTree(&Node{Value: 1})
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	tree.PrintTree(tree.Root)
}
