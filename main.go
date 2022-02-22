package main

import (
	"fmt"
	"os"
	"time"

	hello "github.com/ichang0301/learn-golang/1_hello-world"
	integer "github.com/ichang0301/learn-golang/2_integer"
	iteration "github.com/ichang0301/learn-golang/3_iteration"
	arrays "github.com/ichang0301/learn-golang/4_arrays-and-slices"
	structs "github.com/ichang0301/learn-golang/5_structs_methods_interfaces"
	pointer "github.com/ichang0301/learn-golang/6_pointers-and-errors"
	maps "github.com/ichang0301/learn-golang/7_maps"
	dependancy_injection "github.com/ichang0301/learn-golang/8_dependancy-injection"
	mock "github.com/ichang0301/learn-golang/9_mocking"
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
	if err := oh_wallet.Withdraw(20); err != nil {
		fmt.Errorf("Failed to Withdraw(): %q\n", err)
	}
	fmt.Println(oh_wallet.Balance())

	var dictionary = maps.Dictionary{}
	if err := dictionary.Add("test", "this is just a test"); err != nil {
		fmt.Errorf("Failed to Add() to dictionary: %q\n", err)
	}
	fmt.Println(dictionary)
	if definition, err := dictionary.Search("test"); err == nil {
		fmt.Printf("definition: %s\n", definition)
	}
	if err := dictionary.Update("test", "new definition"); err != nil {
		fmt.Errorf("Failed to Update() to dictionary: %q\n", err)
	}
	fmt.Println(dictionary)
	dictionary.Delete("test")
	fmt.Println(dictionary)

	dependancy_injection.Greet(os.Stdout, "Mike")

	sleeper := &mock.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mock.Countdown(os.Stdout, sleeper)
}
