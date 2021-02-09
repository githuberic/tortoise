package law_of_demeter

import "fmt"

type GoodTeamLeader struct {
	iID int
	sName string
}

func NewGoodTeamLeader(id int, name string) *GoodTeamLeader {
	return &GoodTeamLeader{
		id,
		name,
	}
}

func (p *GoodTeamLeader) CountOpeningTasks() int {
	member := NewGoodTeamMember(11, "王Member")
	sum := member.countOpeningTasks()

	fmt.Printf("%v CountOpeningTasks, got %v", p.sName, sum)
	return sum
}
