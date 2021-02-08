package law_of_demeter

type BadTeamMember struct {
	iID   int
	sName string
}

func NewBadTeamMember(id int, name string) *BadTeamMember {
	return &BadTeamMember{
		id,
		name,
	}
}

func (me *BadTeamMember) countOpeningTasks(lstTasks []*Task) int {
	sum := 0
	for _, it := range lstTasks {
		if it.Status() == OPENING {
			sum++
		}
	}
	return sum
}
