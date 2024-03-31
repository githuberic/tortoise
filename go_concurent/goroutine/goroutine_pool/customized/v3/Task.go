package v3


// 定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error
}

//通过NewTask来创建一个Task
func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

// 执行Task任务的方法
func (t *Task) Execute() {
	// 调用任务所绑定的函数
	t.f()
}
