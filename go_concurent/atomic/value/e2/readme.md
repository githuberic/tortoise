# atomic.Value代替sync.RWMutex

下面是一个常见的使用场景：应用程序定期的从外界获取最新的配置信息，然后更改自己内存中维护的配置变量。

工作线程根据最新的配置来处理请求。

### 使用sync.RWMutex读写配置

发现好些使用sync.RWMutext的使用场景：项目启动时候对高频数据缓存到内存缓存中，同时每隔一段时间重新写一下这个缓存（用一个全局变量）：

```go

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

```

看到以上场景是一个协程写，接口并发读，而写的过程是直接修改变量的值(切片引用类型，修改了全局变量指向的底层数组)，只是这一个写的过程，

却要在每次读的时候加读锁，其实只要换成保证原子操作的变量赋值和读取就行了，使用atomic.Value

atomic.Value封装了一个interface{}类型的v变量，简单粗暴，同时提供了Load和Store原子操作，

用于给v原子赋值和读取，简单修改下：

```go
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
```

原子操作由底层硬件支持，而锁则由操作系统的调度器实现。

锁应当用来保护一段逻辑，对于一个变量更新的保护，原子操作通常会更有效率