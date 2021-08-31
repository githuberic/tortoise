package basic_queue

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	headNode *Node
	tailNode *Node
	length   int
	count    int
}

func NewQueue(length int) *LinkedList {
	return &LinkedList{length: length}
}

// 入队
func (l *LinkedList) Push(data int) bool {
	if l.count == l.length {
		l.tailNode = nil
		return false
	}

	node := &Node{data: data}
	if l.headNode == nil {
		l.headNode = node
	} else {
		l.tailNode.next = node
	}
	l.tailNode = node
	l.count++

	return true
}

// 循环队列入队
func (l *LinkedList) LoopPush(data int) bool {
	if l.count == l.length {
		l.tailNode = nil
		return false
	}

	node := &Node{data: data, next: l.headNode}
	if l.headNode == nil {
		l.headNode = node
	} else {
		l.tailNode.next = node
	}
	l.tailNode = node
	l.count++

	return true
}

// 出对
func (l *LinkedList) Pop() int {
	if l.count == 0 {
		return 0
	}
	res := l.headNode.data
	l.headNode = l.headNode.next
	l.count--
	return res
}
