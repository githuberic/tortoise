package binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct{ pre, in, post []interface{} }{
	{
		[]interface{}{1, 2, 3},
		[]interface{}{1, 3, 2},
		[]interface{}{3, 2, 1},
	},

	{
		[]interface{}{1, 2, 4, 5, 3, 6, 7},
		[]interface{}{4, 2, 5, 1, 6, 3, 7},
		[]interface{}{4, 5, 2, 6, 7, 3, 1},
	},
	// 可以有多个 testCase
}

func PreIn2Tree(pre, in []interface{}) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

/*
var root TreeNode


func init() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Right.Right = NewNode(5)
}*/

func indexOf(val interface{}, nums []interface{}) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func Test_preOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.pre, preOrderTraversal(root), "输入:%v", tc)
	}
}

func Test_inOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.in, inOrderTraversal(root), "输入:%v", tc)
	}
}

func Test_postOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.post, postOrderTraversal(root), "输入:%v", tc)
	}
}


// 测试
func myTest(s []interface{}) [][]int {
	convertor := &Convertor{}
	root := convertor.ConvertSlice2Tree(s)
	pre := preOrderTraversal(root)
	in := inOrderTraversal(root)
	post := postOrderTraversal(root)
	return [][]int{pre, in, post}
}


