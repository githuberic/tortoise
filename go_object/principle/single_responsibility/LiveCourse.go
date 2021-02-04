package single_responsibility

import "fmt"

type LiveCourse struct {
	CourseInfo
}

func NewLiveCourse(id int, name string) IGoodCourse {
	return &LiveCourse{
		CourseInfo{
			iID: id,
			sName: name,
		},
	}
}
func (me *LiveCourse) Controller() IPlayControl {
	return me
}
func (me *LiveCourse) Play() {
	fmt.Printf("%v play\n", me.Name())
}
func (me *LiveCourse) Pause() {
	fmt.Printf("%v pause\n", me.Name())
}
