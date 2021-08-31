package circle_queue

type CircularQueue struct {
	items         []string
	head, tail, n int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]string, capacity),
		n:     capacity,
	}
}

func (q *CircularQueue) Enqueue(item string) bool {
	if (q.tail+1)%q.n == q.head {
		// full
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.n
	return true
}

func (q *CircularQueue) Dequeue() (string, bool) {
	if q.head == q.tail {
		return "", false
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.n
	return item, true
}
