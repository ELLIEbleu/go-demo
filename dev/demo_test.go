package dev

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestEventToValue(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("break now")
				break
			}
			fmt.Println("inside the select: ")
		}
		fmt.Println("inside the for: ")
	}
}

func TestGen(t *testing.T) {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// ……
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			time.Sleep(time.Second)
			fmt.Println("*************")
		}
	}()
	return ch
}

//
func gen2(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}

func TestGen2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响

	for n := range gen2(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	// ……
}

func TestCopy(t *testing.T) {
	var arr = make([]int32, 10)

	for i := 0; i < 3; i++ {
		brr := []int32{1, 2, 3}

		copy(arr, brr)
	}

	for _, r := range arr {
		fmt.Println(r)
	}
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprint(int(m))
}

func TestB(t *testing.T) {
	var a MyInt
	a = 10
	fmt.Println(a.String())
	//todo
}

func TestBigInt(t *testing.T) {
	bigA := big.NewInt(math.MaxInt64)
	fmt.Printf("BigA Int:%v\n", bigA)
	bigB := big.NewInt(math.MaxInt64)
	fmt.Printf("BigB Int:%v\n", bigB)

	bigC := big.NewInt(math.MaxInt64)
	bigC.Mul(bigA, bigB)
	fmt.Printf("BigC Int:%v\n", bigC)

}

func Test_Arr(t *testing.T) {
	// arr := [2][7]byte{{'A', 'B', 'C', 'B', 'D', 'A', 'B'}, {'B', 'D', 'C', 'A', 'B', 'A'}}
	// for i := range arr {
	// 	for j := range arr[i] {
	// 		fmt.Printf("%q", arr[i][j])
	// 	}
	// 	fmt.Println()
	// }

	x := "ABCBDAB"
	y := "BDCABA"

	m := len(x)
	n := len(y)
	lookup := make([][]int, m+1)
	for i := range lookup {
		lookup[i] = make([]int, n+1)
	}

}

func Test_LCS(t *testing.T) {
	arr := [7]byte{'A', 'B', 'C', 'B', 'D', 'A', 'B'}
	brr := [6]byte{'B', 'D', 'C', 'A', 'B', 'A'}

	crr := buildCrr(arr, brr)
	printCrr(crr)

	findLCS(crr, brr)
}

func printCrr(crr [7][8]int) {
	for _, x := range crr {
		for _, y := range x {
			fmt.Printf("%v  ", y)
		}
		fmt.Println("")
	}
}

func findLCS(crr [7][8]int, brr [6]byte) {
	//逆推最长子序
	// coll := list.New()
	res := []string{}

	for i := len(crr) - 1; i >= 1; i-- {
		start := crr[i][7]
		str := string(brr[i-1])
		j := 7
		x := i - 1

		for {
			if x < 1 {
				break
			}
			if start-crr[i-1][j-1] != 1 {
				break
			}
			str += " "
			str += string(brr[x-1])
			j = j - 1
			x = x - 1
		}
		// coll.PushBack(str)
		res = append(res, str)
	}
	// for val := coll.Front(); val != nil; val = val.Next() {
	// 	fmt.Println(val)
	// }
	//
	for i := len(res) - 1; i > 0; i-- {
		trr := strings.Split(res[i], " ")
		sort.Reverse(sort.StringSlice(trr))
		fmt.Println(strings.Join(trr, " "))
	}
}

func buildCrr(arr [7]byte, brr [6]byte) [7][8]int {
	//crr :=make([][]byte,8)   //初始化 nil
	//crr := [8][7]byte {{}}
	var crr [7][8]int
	// for j := 1; j < 8; j++ {
	// 	crr[0][j] = string(arr[j-1])
	// }
	// for i := 1; i < 7; i++ {
	// 	crr[i][0] = string(brr[i-1])
	// }

	for i, valB := range brr {
		cx := i + 1
		for j, valA := range arr {
			cy := j + 1
			if i == 0 && j == 0 {
				crr[1][1] = compareSame(valA, valB)
			} else {
				xPreVal := crr[cx-1][cy]
				yPreVal := crr[cx][cy-1]
				diagonalVal := compareSame(valA, valB)
				diagonalVal = crr[cx-1][cy-1] + diagonalVal
				var max int
				if xPreVal > yPreVal {
					max = xPreVal
				} else {
					max = yPreVal
				}
				if diagonalVal > max {
					max = diagonalVal
				}

				crr[cx][cy] = max
			}

		}
	}
	return crr
}

func compareSame(valA byte, valB byte) int {
	if valA == valB {
		return 1
	}
	return 0

}

func Test_Reverse(t *testing.T) {
	str := "A B A B C"

	arr := strings.Split(str, " ")
	// sort.Sort(sort.Reverse(sort.StringSlice(arr)))    //逆序排序
	// fmt.Printf("%v", arr)
	var reverseStr []string
	for i := len(arr) - 1; i > 0; i-- {
		reverseStr = append(reverseStr, arr[i])
	}
	fmt.Println(strings.Join(reverseStr, " "))

}

func Test_Analyze(t *testing.T) {

	//Text
}

func Test_Func(t *testing.T) {
	go func() {
		for {
			testPrint()
			fmt.Println("try to re-consume kings landing in 5 seconds")
			time.Sleep(time.Second * 1)
			fmt.Println("start reconnect...")
		}
	}()
}

func testPrint() {
	fmt.Println("print in func test")
}

func Test_Select(t *testing.T) {
	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	//for {
	//	select {
	//	case msg1 := <-ch:
	//		fmt.Println("received: " + msg1)
	//
	//	case <-time.After(time.Minute * 3):
	//		fmt.Println("after 3 minute")
	//		return
	//
	//	}
	//}

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch:
			fmt.Println("received: " + msg1)
		case msg2 := <-ch2:
			fmt.Println("received: " + msg2)

		}
	}
}

func Test_B(t *testing.T) {
	var arr []int
	for i := 0; i < 15; i++ {
		arr = append(arr, i)
	}
	//for i := 0; i <len(arr); i++ {
	//	val := arr[i]
	//	go func() {
	//		fmt.Println(val)
	//	}()
	//}

	//for i := range arr{
	//	val := arr[i]
	//	go func() {
	//		fmt.Println(val)
	//	}()
	//}

	var wg sync.WaitGroup
	//wg.Add(3)
	//go func() {
	//	for i :=0; i< 5 ;i++  {
	//		fmt.Println(arr[i])
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for i :=5;i<10 ; i++{
	//		fmt.Println(arr[i])
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for i :=10; i<15 ; i++  {
	//		fmt.Println(arr[i])
	//	}
	//	wg.Done()
	//}()

	duration := 5
	delta := len(arr)/duration + 1

	for i := 0; i < delta; i++ {
		wg.Add(1)
		start := i * duration
		max := start + duration
		go func() {
			if max <= len(arr) {
				for j := start; j < max; j++ {
					fmt.Println(arr[j])
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestChannel(t *testing.T)  {
	ch := make(chan struct{})
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second *5)
		ch <- struct{}{}
	}()
	data :=<- ch
	fmt.Println(data)
	fmt.Println("finished")
}

/**
 测试阻塞
 */
func TestRequestWork(t *testing.T)  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响

	ch := make(chan struct{})

	RequestWork(ctx,ch)
}


