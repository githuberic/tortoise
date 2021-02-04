package dependence_inversion

import (
	"fmt"
	"testing"
)

type ICourse interface {
	ID() int
	Name() string
	SetUser(IUser)
	Study()
}

type IUser interface {
	ID() int
	Name() string
	Study(ICourse)
}

type GoodUser struct {
	iID int
	sName string
}
func NewGoodUser(id int, name string) IUser {
	return &GoodUser{
		iID: id,
		sName: name,
	}
}

func (me *GoodUser) ID() int {
	return me.iID
}

func (me *GoodUser) Name() string {
	return me.sName
}
func (me *GoodUser) Study(course ICourse) {
	course.SetUser(me)
	course.Study()
}


type GolangCourse struct {
	iID int
	sName string
	xCurrentUser IUser
}
func NewGolangCourse() ICourse {
	return &GolangCourse{
		iID: 11,
		sName: "golang",
		xCurrentUser: nil,
	}
}
func (me *GolangCourse) ID() int {
	return me.iID
}
func (me *GolangCourse) Name() string {
	return me.sName
}
func (me *GolangCourse) SetUser(user IUser) {
	me.xCurrentUser = user
}
func (me *GolangCourse) Study() {
	fmt.Printf("%v is learning %v\n", me.xCurrentUser.Name(), me.Name())
}

func TestVerify(t *testing.T)  {
	gu := NewGoodUser(2,"lgq")
	gu.Study(NewGolangCourse())
}
// https://studygolang.com/articles/33101