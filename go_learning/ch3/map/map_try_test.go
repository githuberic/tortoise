package _map

import "testing"

func TestIntMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}



func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(opt int) int{}

	m[1] = func(opt int) int { return opt }
	m[2] = func(opt int) int { return opt * opt }
	m[3] = func(opt int) int { return opt * opt * opt }

	t.Log(m[1](2), m[2](2), m[3](2))
}


func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	// 当访问的key不存在时，仍会返回0值
	if v, ok := m1[4]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("ok value=",ok)
		t.Log("key 3 is not existing.")
	}
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))

	// 删除元素
	delete(mySet, 1)
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

func TestTravelMap(t *testing.T)  {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
