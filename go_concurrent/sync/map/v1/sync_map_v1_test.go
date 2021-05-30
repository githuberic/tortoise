package v1

import (
	"log"
	"sync"
	"testing"
)

func TestVerify(T *testing.T) {
	var m sync.Map

	// 1. 写入
	m.Store("qcrao", 18)
	m.Store("stefno", 20)

	// 2. 读取
	age, _ := m.Load("qcrao")
	log.Print(age.(int))

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		log.Println("name=", name, ",age=", age)
		return true
	})

	// 4. 删除
	m.Delete("qcrao")
	age, ok := m.Load("qcrao")
	log.Println(age,ok)

	// 5. 读取或写入
	m.LoadOrStore("stefno",99)
	age, _ = m.Load("stefno")
	log.Println(age)
}

