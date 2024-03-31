package bfs_v2

type Queue interface {
	Offer(e interface{})
	Poll() interface{}
	Clear() bool
	Size() int
	IsEmpty() bool
}

type LinkedList struct {
	elements []interface{}
}

func NewQueue() *LinkedList {
	return &LinkedList{}
}

func (queue *LinkedList) Offer(e interface{}) {
	queue.elements = append(queue.elements, e)
}

func (queue *LinkedList) Poll() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	firstElement := queue.elements[0]
	queue.elements = queue.elements[1:]
	return firstElement
}

func (queue *LinkedList) Size() int {
	return len(queue.elements)
}

func (queue *LinkedList) IsEmpty() bool {
	return len(queue.elements) == 0
}

func (queue *LinkedList) Clear() bool {
	if queue.IsEmpty() {
		return false
	}
	for i := 0; i < queue.Size(); i++ {
		queue.elements[i] = nil
	}
	queue.elements = nil
	return true
}
