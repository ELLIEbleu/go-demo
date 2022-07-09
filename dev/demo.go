package dev

import (
	"context"
	"testing"
	"time"
	"unsafe"
)

func test3(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	a := 0
	i := 0

	var tmp int
	p := uintptr(unsafe.Pointer(&l[0]))

	if i >= 5 {
		goto end
	}
body:
	tmp = *(*int)(unsafe.Pointer(p))
	p += unsafe.Sizeof(l[0])
	i++
	a += tmp
	if i < 5 {
		goto body
	}
end:
	println(t)
}

func test(t *testing.T) {
	l := []int{9, 45, 23, 67, 78}
	a := 0

	for _, v := range l {
		a += v
	}

	println(t)
}

func test2(t *testing.T) {
	i := 0
	sum := 0
	goto end
start:
	sum += i
	i++
end:
	if i < 6 {
		goto start
	}

	println(sum)
}

func hardWork(job interface{}) error{
	time.Sleep(time.Minute)
	return nil
}

func RequestWork(ctx context.Context, job interface{}) error {
	return hardWork(job)
}

