package dfs_v2

import "fmt"

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Val: data}
}

func (root *Node) DFS() {
	if root == nil {
		return
	}

	res := NewStack()
	s := NewStack()
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*Node)
		if cur.Left != nil {
			s.Push(cur.Left)
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		//fmt.Println(cur.Val)
		res.Push(cur)
	}

	for !res.Empty() {
		fmt.Printf("%v", res.Pop().(*Node).Val)
	}
}

func (root *Node) PreTravesal() {
	if root == nil {
		return
	}

	s := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		fmt.Println(cur.Val)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

func (root *Node) InTravesal() {
	if root == nil {
		return
	}

	s := NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		if s.Empty() {
			break
		}

		cur = s.Pop().(*Node)
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

func (root *Node) PostTravesal() {
	if root == nil {
		return
	}

	s := NewStack()
	out := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		out.Push(cur)

		if cur.Left != nil {
			s.Push(cur.Left)
		}

		if cur.Right != nil {
			s.Push(cur.Right)
		}
	}

	for !out.Empty() {
		cur := out.Pop().(*Node)
		fmt.Println(cur.Val)
	}
}
