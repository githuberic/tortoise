## 匿名变量

在使用多重赋值时，如果不需要在左值中接收变量，可以使用匿名变量（anonymous variable）。

匿名变量的表现是一个下画线_，使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。例如：

```go
import (
"fmt"
"testing"
)

func GetData(x, y int) (int, int) {
return x, x + y
}

func TestVerifyAV(t *testing.T) {
a, _ := GetData(1, 2)
_, sum := GetData(1, 2)
fmt.Println(a, sum)
}
```

GetData() 是一个函数，拥有两个整型返回值。

代码说明如下：
第 1 次调用 需要获取第一个返回值，所以将第二个返回值的变量设为下画线。
第 2 次调用 行将第一个返回值的变量设为匿名。

匿名变量不占用命名空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。