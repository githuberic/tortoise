package err_wrapper

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {

	/**
	我们需要包装一下错误，而不是干巴巴地把err返回到上层，我们需要把一些执行的上下文加入。
	*/
	if err != nil {
		return fmt.Errorf("something failed: %v", err)
	}
}

/**
Go 语言的开发者中，更为普遍的做法是将错误包装在另一个错误中，同时保留原始内容：
*/
type authorizationError struct {
	operation string
	err       error // original error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}