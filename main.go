package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	hello "github.com/ichang0301/learn-golang/01_hello_world"
	integer "github.com/ichang0301/learn-golang/02_integer"
	iteration "github.com/ichang0301/learn-golang/03_iteration"
	arrays "github.com/ichang0301/learn-golang/04_arrays-and-slices"
	structs "github.com/ichang0301/learn-golang/05_structs_methods_interfaces"
	pointer "github.com/ichang0301/learn-golang/06_pointers_and_errors"
	maps "github.com/ichang0301/learn-golang/07_maps"
	dependancy_injection "github.com/ichang0301/learn-golang/08_dependancy_injection"
	mock "github.com/ichang0301/learn-golang/09_mocking"
	synchronize "github.com/ichang0301/learn-golang/13_sync"
	roman_numeral "github.com/ichang0301/learn-golang/15_roman_numerals"
	clockface_svg "github.com/ichang0301/learn-golang/16_math/svg"
	blogposts "github.com/ichang0301/learn-golang/17_reading_files"
	templating "github.com/ichang0301/learn-golang/18_templating"
)

func main() {
	// 01_hello_world
	fmt.Println(hello.Hello("world", ""))

	// 02_integer
	fmt.Println(integer.Add(1, 1))

	// 03_iteration
	fmt.Println(iteration.Repeat("!", 3))

	// 04_arrays-and-slices
	fmt.Println(arrays.SumArray([5]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumSlices([]int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}))
	fmt.Println(arrays.SumAllTails([]int{0, 9}, []int{1, 2, 3}))

	// 05_structs_methods_interfaces
	fmt.Println(structs.Perimeter(structs.Rectangle{Width: 5.0, Height: 10.0}))
	fmt.Println(structs.Rectangle{Width: 10.0, Height: 15.0}.Area())
	fmt.Println(structs.Circle{Radius: 1.0}.Area())
	fmt.Println(structs.Triangle{Base: 5.0, Height: 4.0}.Area())

	// 06_pointers_and_errors
	oh_wallet := pointer.Wallet{}
	oh_wallet.Deposit(30)
	oh_wallet.Withdraw(20)
	fmt.Println(oh_wallet.Balance())

	// 07_maps
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

	// 08_dependancy_injection
	dependancy_injection.Greet(os.Stdout, "Mike")

	// 09_mocking
	sleeper := &mock.ConfigurableSleeper{Duration: 1 * time.Second, SleepDuration: time.Sleep}
	mock.Countdown(os.Stdout, sleeper)

	// 13_sync
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

	// 15_roman_numerals
	fmt.Println(roman_numeral.ConvertToRoman(0))
	fmt.Println(roman_numeral.ConvertToRoman(1984))
	fmt.Println(roman_numeral.ConvertToRoman(4000))
	fmt.Println(roman_numeral.ConvertToArabic("N"))
	fmt.Println(roman_numeral.ConvertToArabic("MCMLXXXIV"))

	// 16_math
	t := time.Now()
	fmt.Println(t)
	const mathResultDirectoryPath = "16_math/result/"
	if err := os.MkdirAll(mathResultDirectoryPath, 0755); err != nil {
		log.Fatal(err)
	}

	const mathResultFileName = "clockface.svg"
	const mathResultFilePath = mathResultDirectoryPath + mathResultFileName
	if err := os.WriteFile(mathResultFilePath, []byte(""), 0644); err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(mathResultFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	clockface_svg.SVGWriter(f, t)

	// 17_reading_files
	posts, err := blogposts.PostsFromFS(os.DirFS("17_reading_files/input"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)

	// 18_templating
	templatingPosts, err := templating.NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}

	convertedPosts := []templating.Post{
		{
			Title:       posts[0].Title,
			Body:        posts[0].Body,
			Description: posts[0].Description,
			Tags:        posts[0].Tags,
		},
		{
			Title:       posts[1].Title,
			Body:        posts[1].Body,
			Description: posts[1].Description,
			Tags:        posts[1].Tags,
		},
	}
	fmt.Println(templatingPosts.Render(os.Stdout, convertedPosts[0]))
	fmt.Println(templatingPosts.RenderIndex(os.Stdout, convertedPosts))
}
