package dfs_v2

import (
	"testing"
)

/**
Tree深度遍历测试
*/
func TestDFS(t *testing.T) {
	root := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)
	n3 := NewNode(4)
	n4 := NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4

	root.DFS()
}

