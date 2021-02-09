package v2

import "testing"

/**
课程接口测试用例
*/
func TestVerify(t *testing.T) {
	fnShowCourse := func(it ICourse) {
		t.Logf("id=%v, name=%v, price=%v\n", it.ID(), it.Name(), it.Price())
	}

	c1 := NewGolangCourse(1, "golang课程", 100)
	fnShowCourse(c1)

	c2 := NewDiscountGolangCourse(2, "golang优惠课程", 100, 0.6)
	fnShowCourse(c2)
}
