package main

import (
	"fmt"

	hello "github.com/ichang0301/learn-golang/1_hello-world"
	integer "github.com/ichang0301/learn-golang/2_integer"
	iteration "github.com/ichang0301/learn-golang/3_iteration"
	arrays "github.com/ichang0301/learn-golang/4_arrays-and-slices"
)

func main() {
	fmt.Println(hello.Hello("world", ""))

	fmt.Println(integer.Add(1, 1))

	fmt.Println(iteration.Repeat("!", 3))

	fmt.Println(arrays.SumArray([5]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumSlices([]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAllTails([]int{0, 9}, []int{1, 2, 3}))
}
