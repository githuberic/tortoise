
package polymorphism

import (
	"fmt"
	"testing"
)

//Person接口
type Person interface {
	ToSchool()
}

//学生类
type Student struct {
	work string
}
//学生类实现Person接口
func (this *Student) ToSchool() {
	fmt.Println("Student ", this.work)
}

//老师类
type Teacher struct {
	work string
}
//老师类实现Person接口
func (this *Teacher) ToSchool() {
	fmt.Println("Teacher ", this.work)
}

//工厂模式函数，根据传入工作不同动态返回不同类型
func Factory(work string) Person {
	switch work {
	case "study":
		return &Student{work: "study"}
	case "teach":
		return &Teacher{work: "teach"}
	default:
		panic("no such profession")
	}
}

func TestVerifyV1(t *testing.T) {
	person := Factory("study")
	person.ToSchool() //Student  study

	person = Factory("teach")
	person.ToSchool() //Teacher  teach
}

// https://blog.csdn.net/baolingye/article/details/111166714?utm_medium=distribute.pc_relevant.none-task-blog-baidujs_title-7&spm=1001.2101.3001.4242