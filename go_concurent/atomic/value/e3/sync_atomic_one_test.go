package e3

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 暂时废弃
func TestVerifyOne(t *testing.T) {
	type Map map[string]string

	var m atomic.Value
	m.Store(make(Map))

	var mu sync.Mutex // used only by writers
	read := func(key string) string {
		m1 := m.Load().(Map)
		return m1[key]
	}

	insert := func(key, val string) {
		mu.Lock() // 与潜在写入同步
		defer mu.Unlock()

		m1 := m.Load().(Map) // 导入struct当前数据
		m2 := make(Map)      // 创建新值
		for k, v := range m1 {
			m2[k] = v
		}
		m2[key] = val
		m.Store(m2) // 用新的替代当前对象
	}
	_, _ = read, insert
}
