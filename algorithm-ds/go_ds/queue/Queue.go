package queue

type Queue interface {
	EnQueue(v interface{}) bool
	DeQueue() interface{}
}
