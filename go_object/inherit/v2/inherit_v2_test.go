package v2

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age int
}
func (this *Person) WriteCode(){
	fmt.Println("Write code")
}

type Programmer struct {
	skills string
}
func (this *Programmer) WriteCode(){
	fmt.Println("我会设计项目")
}

type ProjectManager struct {
	Person
	Programmer
}

func TestVerifyV2(t *testing.T)  {
	projectManager := ProjectManager {}
	// 因为有所继承的结构体中有两个相同的方法名字，不能通过平常的调用方式
	// projectManager.WriteCode()这样是错误的，编译不过去，分不清楚到底是那个结构体中的方法
	// 所以要这样用结构体的实例.继承的结构体的名字.字段或者方法名
	projectManager.Programmer.WriteCode()
}
