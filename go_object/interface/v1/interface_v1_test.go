package v1

import (
	"fmt"
	"testing"
)

//Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树..")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

// 让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

//让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳..")
}

func TestVerifyV1(t *testing.T) {
	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey {
			Name : "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}

// https://blog.csdn.net/cui_yonghua/article/details/94362727?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control

// https://blog.csdn.net/cui_yonghua/article/details/94410677?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control