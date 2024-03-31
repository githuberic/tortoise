# Set

golang-set -A simple set type for the Go language. Also used by Docker, 1Password, Ethereum.

在github上已经有了一个成熟的包，名为golang-set，包中提供了线程安全和非线程安全的set。提供了五个set函数：

<blockquote>
// NewSet创建并返回空集的引用，结果集上的操作是线程安全的
func NewSet(s ...interface{}) Set {}

// NewSetFromSlice从现有切片创建并返回集合的引用，结果集上的操作是线程安全的
func NewSetFromSlice(s []interface{}) Set {}

// NewSetWith创建并返回具有给定元素的新集合，结果集上的操作是线程安全的
func NewSetWith(elts ...interface{}) Set {}

// NewThreadUnsafeSet创建并返回对空集的引用，结果集上的操作是非线程安全的
func NewThreadUnsafeSet() Set {}

// NewThreadUnsafeSetFromSlice创建并返回对现有切片中集合的引用，结果集上的操作是非线程安全的。
func NewThreadUnsafeSetFromSlice(s []interface{}) Set {}
</blockquote>