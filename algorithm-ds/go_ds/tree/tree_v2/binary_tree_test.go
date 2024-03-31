package tree_v2

import (
	"algorithm-ds/go_ds/tree"
	"testing"
)

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.Root.Left = tree.NewNode(3)
	binaryTree.Root.Right = tree.NewNode(4)
	binaryTree.Root.Right.Left = tree.NewNode(5)

	binaryTree.InOrderTraverse()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.Root.Left = tree.NewNode(3)
	binaryTree.Root.Right = tree.NewNode(4)
	binaryTree.Root.Right.Left = tree.NewNode(5)

	binaryTree.PreOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.Root.Left = tree.NewNode(3)
	binaryTree.Root.Right = tree.NewNode(4)
	binaryTree.Root.Right.Left = tree.NewNode(5)

	binaryTree.PostOrderTraverse()
}
