package go_runtime

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func findRoad(t *int) {
	// 一般在这里进行资源回收
	log.Println("test:", *t)
}

func entry() {
	var rd int = int(1111)
	p := &rd
	// 解除绑定并执行对应函数下一次gc在进行清理
	runtime.SetFinalizer(p, findRoad)
}

func TestFinalizer(t *testing.T) {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 100)
		runtime.GC()
	}
}
