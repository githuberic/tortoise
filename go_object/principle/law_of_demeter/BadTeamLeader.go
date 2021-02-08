package law_of_demeter

import (
	"fmt"
)

type BadTeamLeader struct {
	iID int
	sName string
}

func NewBadTeamLeader(id int, name string) *BadTeamLeader {
	return &BadTeamLeader{
		id,
		name,
	}
}

// 不好的ITeamLeader实现, 同时耦合了Task和BadTeamMember两个类
func (me *BadTeamLeader) CountOpeningTasks() int {
	tasks := LoadTaskList()
	member := NewBadTeamMember(11, "王Member")
	sum := member.countOpeningTasks(tasks)

	fmt.Printf("%v CountOpeningTasks, got %v", me.sName, sum)
	return sum
}

