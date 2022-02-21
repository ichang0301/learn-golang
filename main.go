package main

import (
	"fmt"

	hello "github.com/ichang0301/learn-golang/1_hello-world"
	integer "github.com/ichang0301/learn-golang/2_integer"
	iteration "github.com/ichang0301/learn-golang/3_iteration"
	arrays "github.com/ichang0301/learn-golang/4_arrays-and-slices"
	structs "github.com/ichang0301/learn-golang/5_structs-methods-interfaces"
	pointer "github.com/ichang0301/learn-golang/6_pointers-and-errors"
)

func main() {
	fmt.Println(hello.Hello("world", ""))

	fmt.Println(integer.Add(1, 1))

	fmt.Println(iteration.Repeat("!", 3))

	fmt.Println(arrays.SumArray([5]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumSlices([]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAllTails([]int{0, 9}, []int{1, 2, 3}))

	fmt.Println(structs.Perimeter(structs.Rectangle{Width: 5.0, Height: 10.0}))
	fmt.Println(structs.Rectangle{Width: 10.0, Height: 15.0}.Area())
	fmt.Println(structs.Circle{Radius: 1.0}.Area())
	fmt.Println(structs.Triangle{Base: 5.0, Height: 4.0}.Area())

	oh_wallet := pointer.Wallet{}
	oh_wallet.Deposit(30)
	oh_wallet.Withdraw(20)
	fmt.Println(oh_wallet.Balance())
}
