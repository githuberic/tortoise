package work_manager

import (
	"fmt"
	"runtime"
	"sync"
)

type Worker interface {
	Start()
	Stop()
}

type Master struct {
	sync.WaitGroup
	WorkerList []Worker
}

func Stack() []byte {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

func NewMaster() *Master {
	return &Master{WorkerList: make([]Worker, 0, 10)}
}

func (m *Master) AddWorker(w Worker) {
	m.WorkerList = append(m.WorkerList, w)
}

func (m *Master) Start() {
	m.Add(len(m.WorkerList))

	for _, worker := range m.WorkerList {
		go func(w Worker) {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Printf("WorkerManager error, error:%v, stack:%v\n", err, string(Stack()))
				}
			}()
			w.Start()
		}(worker)
	}
}

func (m *Master) Stop() {
	for _, worker := range m.WorkerList {
		go func(w Worker) {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Printf("WorkerManager error, error:%v, stack:%v\n", err, string(Stack()))
				}
			}()
			worker.Stop()
			m.Done()
		}(worker)
	}
}
