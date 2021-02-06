package liskov_substitution

import (
	"testing"
	"tortoise/go_object/principle/liskov_substitution/bad"
	"tortoise/go_object/principle/liskov_substitution/good"
)

func TestBad(t *testing.T) {
	fnCallAndLog := func(fn func() error) {
		err := fn()
		if err != nil {
			t.Logf("error = %s", err.Error())
		}
	}

	bb := bad.NewBadNormalBird(1, "正常鸟(非鸵鸟)")
	fnCallAndLog(bb.Tweet)
	fnCallAndLog(bb.Fly)

	bo := bad.NewBadOstrichBird(2, "鸵鸟")
	fnCallAndLog(bo.Tweet)
	fnCallAndLog(bo.Fly)
	if it, ok := bo.(*bad.BadOstrichBird); ok {
		fnCallAndLog(it.Run)
	}
}

func TestGood(t *testing.T) {
	fnCallAndLog := func(fn func() error) {
		err := fn()
		if err != nil {
			t.Logf("error = %s", err.Error())
		}
	}

	fnTestGoodBird := func(gb good.IGoodBird) {
		fnCallAndLog(gb.Tweet)
		if it, ok := gb.(good.IFlyableBird);ok {
			fnCallAndLog(it.Fly)
		}
		if it, ok := gb.(good.IRunnableBird);ok {
			fnCallAndLog(it.Run)
		}
	}

	fnTestGoodBird(good.NewGoodFlyableBird(11, "飞鸟"))
	fnTestGoodBird(good.NewGoodOstrichBird(12, "鸵鸟"))
}

// from https://studygolang.com/articles/33105
