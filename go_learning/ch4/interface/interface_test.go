package interface_test

import "testing"

// 定义接口
type Programmer interface {
	WriteHelloWorld() string
}


type GoProgrammer struct {
}

// duck type(鸭子类型) Programmer=WriteHelloWorld() string
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	// 类型
	// GoProgrammer 实现了Programmer的接口,无依赖绑定 impliement 只要方法签名一样即可
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
