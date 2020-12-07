package main

import "fmt"

/**
这里return先给result赋值为6，之后执行defer，result变为42，最后返回42
*/
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

/**
这里return确定返回值6，之后defer修改result，最后函数返回return确定的返回值
*/
func f2() int {
	result := 6
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
		fmt.Printf("result=%d\n",result)
	}()
	return result
}

func multiDefer() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}

	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func main() {
	fmt.Println(f())
	fmt.Println(f2())
	multiDefer()

	// 由于defer这里调用的func没有参数，等执行的时候，i已经为0，因而这里输出3个0
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}
