package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	// slice
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// type、length、capacity
	s2 := make([]int, 3, 5)
	t.Log("length=", len(s2), "capacity=", cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log("length=", len(s2), "capacity=", cap(s2))
}

// 切片可变长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// capacity 每次都是2倍增幅
		s = append(s, i)
		t.Log("length=", len(s), "capacity=", cap(s))
	}
}

// 切片共享存储结构， slice 容量可以伸缩，切片不可比较
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	/*
		if a == b { //切片只能和nil比较
			t.Log("equal")
		}*/
	t.Log(a, b)
}
