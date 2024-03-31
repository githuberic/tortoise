package _map

import (
	"log"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkMap(t *testing.B) {
	var mMap = &MutexMap{
		Map: make(map[string]interface{}),
	}

	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		mMap.writeMap(key, i)

		value, ok := mMap.readMap(key)
		if ok {
			log.Print(value)
		}
	}
}

func BenchmarkMutexMap(t *testing.B) {
	var mMap = &MutexMap{
		Map: make(map[string]interface{}),
	}

	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i)

		go func(key string, v int) {
			mMap.writeMap(key, v)
		}(key, i)

		go func(key string) {
			value, ok := mMap.readMap(key)
			if ok {
				log.Print(value)
			}
		}(key)
	}
}

func BenchmarkRWMutexMap(b *testing.B) {
	var mMap = &RWMutexMap{
		Map: make(map[string]interface{}),
	}

	for i := 1; i < 1000; i++ {
		key := strconv.Itoa(i)
		mMap.Set(key, i)

		value, ok := mMap.Get(key)
		if ok {
			log.Println(value)
		}
		/*
			key := strconv.Itoa(i)
			go func() {
				mMap.Set(key, i)
			}()

			go func() {
				value, ok := mMap.Get(key)
				if ok {
					log.Println(value)
				}
			}()*/
	}
}

func BenchmarkChMap(t *testing.B) {
	mMap := NewChannelMap()

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.add(i, i)
		}()

		go func() {
			value := mMap.find(i)
			if value > 0 {
				log.Printf("Value=%d", value)
			}
		}()
	}
}

func BenchmarkChPoolMap(t *testing.B) {
	mMap := NewChannelPoolMap()

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.add(i, i)
		}()

		go func() {
			value := mMap.find(i)
			if value != nil && *value > 0 {
				log.Print(*value)
			}
		}()
	}
}

func BenchmarkSyncMap(t *testing.B) {
	var m sync.Map

	for i := 1; i < 1000; i++ {
		go func() {
			m.Store(i, i)
		}()

		go func() {
			// loop
			m.Range(func(key, value interface{}) bool {
				log.Println("key=", key.(int), ",value=", value.(int))
				return true
			})
		}()
	}
}
