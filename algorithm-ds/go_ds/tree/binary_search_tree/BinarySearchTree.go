package binary_search_tree

import (
	"algorithm-ds/go_ds/tree"
	"algorithm-ds/go_ds/tree/tree_v2"
)

type BST struct {
	*tree_v2.BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{BinaryTree: tree_v2.NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BST) Find(v interface{}) *tree.Node {
	p := this.Root
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 { //v > nodeV
			p = p.Right
		} else { //v < nodeV
			p = p.Left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.Root
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if nil == p.Right {
				p.Right = tree.NewNode(v)
				break
			}
			p = p.Right
		} else {
			if nil == p.Left {
				p.Left = tree.NewNode(v)
				break
			}
			p = p.Left
		}
	}
	return true
}

func (this *BST) Delete(v interface{}) bool {
	var pp *tree.Node = nil
	p := this.Root
	deleteLeft := false
	for nil != p {
		compareResult := this.compareFunc(v, p.Val)
		if compareResult > 0 {
			pp = p
			p = p.Right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.Left
			deleteLeft = true
		} else {
			break
		}
	}

	if nil == p { //需要删除的节点不存在
		return false
	} else if nil == p.Left && nil == p.Right { //删除的是一个叶子节点
		if nil != pp {
			if deleteLeft {
				pp.Left = nil
			} else {
				pp.Right = nil
			}
		} else { //根节点
			this.Root = nil
		}
	} else if nil != p.Right { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p
		q := p.Right //向右走一步
		fromRight := true
		for nil != q.Left { //向左走到底
			pq = q
			q = q.Left
			fromRight = false
		}
		if fromRight {
			pq.Right = nil
		} else {
			pq.Left = nil
		}
		q.Left = p.Left
		q.Right = p.Right
		if nil == pp { //根节点被删除
			this.Root = q
		} else {
			if deleteLeft {
				pq.Left = q
			} else {
				pq.Right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.Left = p.Left
			} else {
				pp.Right = p.Left
			}
		} else {
			if deleteLeft {
				this.Root = p.Left
			} else {
				this.Root = p.Left
			}
		}
	}
	return true
}


func (this *BST) Min() *tree.Node {
	p := this.Root
	for nil != p.Left {
		p = p.Left
	}
	return p
}

func (this *BST) Max() *tree.Node {
	p := this.Root
	for nil != p.Right {
		p = p.Right
	}
	return p
}