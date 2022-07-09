package dev

import (
	"fmt"
	"testing"
	"time"
)

type Person struct {
	Age int
}

func ReturnSliceWithPointers(size int) []* Person{
	res := make([]* Person, size)

	for i :=0; i< size ;i++{
		res[i] = &Person{}
	}

	return  res
}

func ReturnSliceWithStructs(size int) []Person{
	res := make([]Person,size)

	for i :=0; i< size; i++{
		res[i] = Person{}
	}

	return res
}

func Benchmark_ReturnSliceWithPointers( b *testing.B)  {
	for i :=0; i< b.N;i++{
		ReturnSliceWithPointers(10000)
	}
}

func Benchmark_ReturnSliceWithStructs(b *testing.B)  {
	for i :=0; i< b.N; i++{
		ReturnSliceWithStructs(10000)
	}
}

func TestSliceWithPointers(t *testing.T){
	in := []int{1,2,3}

	var out [] *int
	var ret [] int

	for _,i := range in {
		i := i               //迭代器变量复制到新的变量
		out = append(out,&i)
		ret = append(ret,i)
	}

	fmt.Println("out values:",*out[0],*out[1],*out[2])
	fmt.Println("out address:",out[0],out[1],out[2])
	fmt.Println("ret values:",ret[0],ret[1],ret[2])
	fmt.Println("ret address:",&ret[0],&ret[1],&ret[2])


	list := []int{1,2,3}
	for _, v := range  list{
		v := v
		go func() {
			fmt.Printf("%d ", v)
		}()
	}

	time.Sleep(time.Second * 20)
}
