package basic

import (
	"fmt"
	"testing"
)

type mapKey struct {
	key int
}

func TestB1(t *testing.T) {
	var m = make(map[mapKey]string)

	var key = mapKey{10}

	m[key] = "hello"
	fmt.Printf("m[key]=%s\n", m[key])

	// 修改key的字段的值后再次查询map，无法获取刚才add进去的值
	key.key = 100
	fmt.Printf("再次查询m[key]=%s\n", m[key])
}

/***
Go 内建的 map 对象不是线程（goroutine）安全的，并发读写的时候运行时会有检查，遇到并发问题就会导致 panic。
*/
func TestB2(t *testing.T) {
	var m = make(map[int]int, 10)

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[2]
		}
	}()

	select {}
}
