package single_responsibility

import "testing"

func Test_SimpleResponsibility(t *testing.T) {
	fnTestGoodCourse := func(gc IGoodCourse) {
		pc := gc.Controller()
		pc.Play()
		pc.Pause()
		if rc, ok := pc.(IReplayControl);ok {
			rc.Forward(30)
			rc.Backward(30)
		}
	}

	fnTestGoodCourse(NewLiveCourse(11, "直播课"))
	fnTestGoodCourse(NewReplayCourse(12, "录播课"))
}
// from https://studygolang.com/articles/33102
