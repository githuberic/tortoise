package _map

import (
	"log"
	"sync"
	"testing"
)

func TestVerify(T *testing.T) {
	var m sync.Map

	// 1. 写入
	m.Store("lgq", 18)
	m.Store("ldx", 20)
	m.Store("lyy", 20)

	// 2. 读取
	age, _ := m.Load("lgq")
	log.Print(age.(int))

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		log.Println("name=", name, ",age=", age)
		return true
	})

	// 4. 删除
	m.Delete("lyy")
	age, ok := m.Load("lgq")
	log.Println(age,ok)

	// 5. 读取或写入
	m.LoadOrStore("ldx",99)
	age, _ = m.Load("ldx")
	log.Println(age)
}

