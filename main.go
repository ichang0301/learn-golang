package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	hello "github.com/ichang0301/learn-golang/01_hello-world"
	integer "github.com/ichang0301/learn-golang/02_integer"
	iteration "github.com/ichang0301/learn-golang/03_iteration"
	arrays "github.com/ichang0301/learn-golang/04_arrays-and-slices"
	structs "github.com/ichang0301/learn-golang/05_structs_methods_interfaces"
	pointer "github.com/ichang0301/learn-golang/06_pointers-and-errors"
	maps "github.com/ichang0301/learn-golang/07_maps"
	dependancy_injection "github.com/ichang0301/learn-golang/08_dependancy-injection"
	mock "github.com/ichang0301/learn-golang/09_mocking"
	synchronize "github.com/ichang0301/learn-golang/13_sync"
	roman_numeral "github.com/ichang0301/learn-golang/15_roman-numerals"
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

	var dictionary = maps.Dictionary{}
	dictionary.Add("test", "this is just a test")
	fmt.Println(dictionary)
	if definition, err := dictionary.Search("test"); err == nil {
		fmt.Printf("definition: %s\n", definition)
	}
	dictionary.Update("test", "new definition")
	fmt.Println(dictionary)
	dictionary.Delete("test")
	fmt.Println(dictionary)

	dependancy_injection.Greet(os.Stdout, "Mike")

	sleeper := &mock.ConfigurableSleeper{Duration: 1 * time.Second, SleepDuration: time.Sleep}
	mock.Countdown(os.Stdout, sleeper)

	const wantedCount int = 100
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	counter := synchronize.NewCounter()
	for i := 0; i < wantedCount; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())

	fmt.Println(roman_numeral.ConvertToRoman(1984))
}
