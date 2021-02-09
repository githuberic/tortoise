package main

import (
	"fmt"
	"testing"
)

import (
	lod "tortoise/go_object/principle/law_of_demeter"
)

func Test_LOD(t *testing.T) {
	bl := lod.NewBadTeamLeader(1, "张Leader")
	bl.CountOpeningTasks()
	fmt.Print("\n")

	gl := lod.NewGoodTeamLeader(2, "李Leader")
	gl.CountOpeningTasks()
	fmt.Print("\n")
}

//from https://studygolang.com/articles/33104
