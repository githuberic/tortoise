package v2

import (
	"github.com/pkg/errors"
	"testing"
)

/**
Go 语言的开发者中，更为普遍的做法是将错误包装在另一个错误中，同时保留原始内容：
*/
type authorizationError struct {
	operation string
	err       error // original error
}

type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}

func TestVerify(t *testing.T) {
	//错误包装
	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	// Cause接口
	switch err := errors.Cause(err).(
	type
	) {
	case *MyError:
		// handle specifically
	default:
		// unknown error
	}
}
