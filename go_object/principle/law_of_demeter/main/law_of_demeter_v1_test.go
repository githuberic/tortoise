package main

import (
	"testing"
)

import (
	lod "tortoise/go_object/principle/law_of_demeter"
)

func Test_LOD(t *testing.T) {
	bl := lod.NewBadTeamLeader(1, "张Leader")
	bl.CountOpeningTasks()

	gl := lod.NewGoodTeamLeader(2, "李Leader")
	gl.CountOpeningTasks()
}

//from https://studygolang.com/articles/33104
