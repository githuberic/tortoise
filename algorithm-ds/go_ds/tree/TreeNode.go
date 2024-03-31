package tree

import "fmt"

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Val: data}
}

func (this *Node) String() string {
	return fmt.Sprintf("v:%+v, Left:%+v, Right:%+v", this.Val, this.Left, this.Right)
}

