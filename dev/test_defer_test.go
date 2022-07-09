package dev

import (
	"fmt"
	"testing"
)

func Test_Defer(t *testing.T) {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
	fmt.Println("************")
	fmt.Println("b return:", b()) // 打印结果为 b return: 2

}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}
