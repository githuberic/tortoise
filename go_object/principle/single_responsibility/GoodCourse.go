package single_responsibility

type IGoodCourse interface {
	ID() int
	Name() string
	Controller() IPlayControl
}

type IPlayControl interface {
	Play()
	Pause()
}

type IReplayControl interface {
	IPlayControl
	Forward(seconds int)
	Backward(seconds int)
}

type CourseInfo struct {
	iID int
	sName string
}
func (me *CourseInfo) ID() int {
	return me.iID
}
func (me *CourseInfo) Name() string {
	return me.sName
}
