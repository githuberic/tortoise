package go_bad

import (
	"io/ioutil"
	"log"
	"testing"
)

// 切片会导致整个底层数组被锁定，底层数组无法释放内存。如果底层数组较大会对内存产生很大的压力。
func TestSliceCopy(t *testing.T) {
	headerMap := make(map[string][]byte)
	for i := 0; i < 5; i++ {
		name := "/path/to/file"
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		headerMap[name] = data[:1]
	}
}
