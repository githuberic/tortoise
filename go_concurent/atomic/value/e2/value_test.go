package e2

import (
	"sync"
	"sync/atomic"
	"time"
)

type cosCred struct {
	Cred []int64
	sync.RWMutex
}

var CosCred *cosCred

// 每分钟写一次
func InitCosCred() {
	CosCred = new(cosCred)
	CosCred.Cred, _ = GetGlobalCredData()
	go func() {
		for range time.NewTicker(10 * time.Minute).C {
			cred, err := GetGlobalCredData()
			if err != nil {
				continue
			}
			CosCred.Lock()
			CosCred.Cred = cred
			CosCred.Unlock()
		}
	}()
}

func GetGlobalCredData() ([]int64, error) {
	return []int64{12312, 312312, 31212}, nil
}

// 接口会并发读
func GetCosCred() []int64 {
	CosCred.RLock()
	defer CosCred.RUnlock()
	if CosCred.Cred == nil {
		return nil
	}
	resp := CosCred.Cred
	return resp
}

var vCosCred atomic.Value

func InitCosCredV() {
	data, err := GetGlobalCredData()
	if err == nil {
		vCosCred.Store(data)
	}

	go func() {
		for range time.NewTicker(10 * time.Minute).C {
			data, err := GetGlobalCredData()
			if err != nil {
				continue
			}
			vCosCred.Store(data)
		}
	}()
}

func GetCosCredV() []int64 {
	resp, _ := vCosCred.Load().([]int64)
	return resp
}
