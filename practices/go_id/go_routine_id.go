package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	go func() {
		gid := GetGoid()
		fmt.Printf("child goruntine1 gid:%v \\n", gid)
	}()
	go func() {
		gid := GetGoid()
		fmt.Printf("child goruntine2 gid:%v \\n", gid)
	}()
	go func() {
		gid := GetGoid()
		fmt.Printf("child goruntine3 gid:%v \\n", gid)
	}()
	go func() {
		gid := GetGoid()
		fmt.Printf("child goruntine4 gid:%v \\n", gid)
	}()
	go func() {
		gid := GetGoid()
		fmt.Printf("child goruntine5 gid:%v \\n", gid)
	}()
	gid := GetGoid()
	fmt.Printf("main goruntine gid:%v \\n", gid)
	time.Sleep(time.Second)
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
