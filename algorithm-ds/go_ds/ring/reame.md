# Ring
Ring是一种循环链表结构，没有头尾，从任意一个节点出发都可以遍历整个链。其定义如下，Value表示当前节点的值：
```go
type Ring struct {
        Value interface{} 
}
```