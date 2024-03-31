package pool

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "lgq", Age: 18})

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func BenchmarkUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B){
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
