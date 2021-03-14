package polymorphism

import (
	"fmt"
	"testing"
)

//员工接口
type Employee interface {
	ToWork()
}

// 程序员类
type Programmer struct {
	work string
}

// 程序员类实现Employee接口
func (p *Programmer) ToWork() {
	fmt.Println("Programmer ", p.work)
}

// 总监类
type Director struct {
	work string
}

// 总监类实现Employee接口
func (p *Director) ToWork() {
	fmt.Println("Director ", p.work)
}

// 老板类
type Boss struct {
	work string
}

// 老板类实现Employee接口
func (p *Boss) ToWork() {
	fmt.Println("Boss ", p.work)
}

//工厂模式函数，根据传入工作不同动态返回不同类型
func Factory(work string) Employee {
	switch work {
	case "code":
		return &Programmer{work: "code"}
	case "pbc":
		return &Director{work: "plan,bpc"}
	case "meeting":
		return &Boss{work: "meeting"}
	default:
		panic("no such profession")
	}
}

func TestVerifyV1(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	person := Factory("code")
	person.ToWork() //Student  study

	person = Factory("meeting")
	person.ToWork() //Teacher  teach
}
