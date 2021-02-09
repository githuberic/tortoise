package law_of_demeter

type GoodTeamMember struct {
	iID int
	sName string
}

func NewGoodTeamMember(id int, name string) *GoodTeamMember {
	return &GoodTeamMember{
		id,
		name,
	}
}

func (p *GoodTeamMember) countOpeningTasks() int {
	sum := 0
	tasks := LoadTaskList()

	for _,it := range tasks {
		if it.Status() == OPENING {
			sum++
		}
	}

	return sum
}