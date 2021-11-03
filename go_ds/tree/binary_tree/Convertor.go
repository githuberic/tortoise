package binary_tree

type Convertor struct{}

func (this *Convertor) ConvertSlice2Tree(arr []interface{}) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}

	queue := make([]*TreeNode, 0, n)
	root := &TreeNode{Val: arr[0]}
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue := queue[1:]

		if i < n && arr[i] != nil {
			left := &TreeNode{Val: arr[i]}
			node.Left = left
			queue = append(queue, left)
		}
		i++

		if i < n && arr[i] != nil{
			right := &TreeNode{Val: arr[i]}
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}
	return root
}
