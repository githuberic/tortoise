package single_responsibility

import "fmt"

type ReplayCourse struct {
	CourseInfo
}

func NewReplayCourse(id int, name string) IGoodCourse {
	return &ReplayCourse{
		CourseInfo{
			iID: id,
			sName: name,
		},
	}
}

func (me *ReplayCourse) Controller() IPlayControl {
	return me
}
func (me *ReplayCourse) Play() {
	fmt.Printf("%v play\n", me.Name())
}
func (me *ReplayCourse) Pause() {
	fmt.Printf("%v pause\n", me.Name())
}
func (me *ReplayCourse) Forward(seconds int) {
	fmt.Printf("%v forward %v\n", me.Name(), seconds)
}
func (me *ReplayCourse) Backward(seconds int) {
	fmt.Printf("%v backward %v\n", me.Name(), seconds)
}