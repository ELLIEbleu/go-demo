package dev

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gonum/matrix/mat64"
)

const (
	mutexLocked = 1 << iota
)

func TestA(t *testing.T) {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//go handle(ctx, 1500*time.Millisecond)
	//
	//select {
	//case <-ctx.Done():
	//	fmt.Println("main", ctx.Err())
	//}
	//context.WithCancel();
	str := string("test")
	fmt.Println(str)

	fmt.Println(mutexLocked)
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestMat(t *testing.T) {
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	ft := mat64.Formatted(a.T(), mat64.Prefix(" "))
	fmt.Println("a^T =\n\n", ft)
}
