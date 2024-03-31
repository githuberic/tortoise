package tree_v1

import (
	"algorithm-ds/go_ds/tree"
	"fmt"
	"testing"
)

func TestDepth(t *testing.T) {
	//root := &Node{1, nil, nil}
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	fmt.Println(Height(root))
}

/**
验证是否满二叉树
*/
func TestLoop(t *testing.T) {
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	values := preOrderTraversal(root)
	fmt.Printf("Pre Order Value %v\n", values)

	inOrderValues := inOrderTraversal(root)
	fmt.Printf("In order Value %v\n", inOrderValues)

	postOrderValues := postOrderTraversal(root)
	fmt.Printf("Post order Value %v\n", postOrderValues)
}

func TestBalanced(t *testing.T) {
	//root := &Node{1, nil, nil}
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = tree.NewNode(31)
	n3.Left = tree.NewNode(41)
	n3.Right = tree.NewNode(42)

	fmt.Println(isBalanced(root))
}

/**
验证是否满二叉树
*/
func TestFull(t *testing.T) {
	//root := &Node{1, nil, nil}
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	n2.Left = tree.NewNode(31)
	n2.Right = tree.NewNode(32)

	fmt.Println(isFull(root))
}

/**
Tree深度遍历测试
*/
func TestDFS(t *testing.T) {
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	n3.Left = tree.NewNode(41)
	n3.Right = tree.NewNode(42)

	values := DFS(root)
	fmt.Printf("DFS Order Value %v\n", values)
}

/**
验证是否满二叉树
*/
func TestBFS(t *testing.T) {
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n4.Left = tree.NewNode(41)
	n4.Right = tree.NewNode(42)

	values := LevelOrderTraversal(root)
	fmt.Printf("BFS Order Value %v\n", values)
}
