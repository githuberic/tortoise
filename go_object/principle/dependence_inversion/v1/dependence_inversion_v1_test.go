package dependence_inversion

import (
	"testing"
	"tortoise/go_object/principle/dependence_inversion/v1/bad"
	"tortoise/go_object/principle/dependence_inversion/v1/good"
)

func TestVerify(t *testing.T) {
	bu := bad.NewBadUser(1, "Tom")
	bu.StudyGolangCourse()

	gu := good.NewGoodUser(2, "lgq")
	gu.Study(good.NewGolangCourse())

	gn := good.NewGoodUser(3, "eric")
	gn.Study(good.NewJavaCourse())
}
