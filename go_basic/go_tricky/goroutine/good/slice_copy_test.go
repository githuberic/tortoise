package go_good

import (
	"io/ioutil"
	"log"
	"testing"
)

// 解决的方法是将结果克隆一份，这样可以释放底层的数组：
func TestSliceCopy(t *testing.T) {
	headerMap := make(map[string][]byte)

	for i := 0; i < 5; i++ {
		name := "/path/to/file"
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		headerMap[name] = append([]byte{}, data[:1]...)
	}
}
