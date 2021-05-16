package interface2_test

// 结构体嵌入
/**
使用了一个 Printable 的接口，而 Country 和 City 都实现了接口方法 PrintStr() 把自己输出
我们可以使用“结构体嵌入”的方式来完成这个事
引入一个叫 WithName的结构体，但是这会带来一个问题：在初始化的时候变得有点乱。
 */
import (
	"fmt"
	"testing"
)

type WithName struct {
	Name string
}

type Country struct {
	WithName
}

type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func TestVerify(t *testing.T)  {
	c1 := Country {WithName{ "China"}}
	c2 := City { WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}


