package law_of_demeter

type TaskStatus int

const OPENING TaskStatus = 0
const DONE TaskStatus = 1
const CANCLED TaskStatus = 2
const DENIED TaskStatus = 3

// 定义任务信息, 以及加载任务清单的方法
type Task struct {
	iID     int
	iStatus TaskStatus
}

func NewTask(id int, status TaskStatus) *Task {
	return &Task{
		id,
		status,
	}
}

func (me *Task) ID() int {
	return me.iID
}

func (me *Task) Status() TaskStatus {
	return me.iStatus
}

func LoadTaskList() []*Task {
	tasks := make([]*Task, 0)
	tasks = append(tasks, NewTask(1, OPENING))
	tasks = append(tasks, NewTask(2, DONE))
	tasks = append(tasks, NewTask(3, CANCLED))
	tasks = append(tasks, NewTask(4, DENIED))
	return tasks
}
