package main

import (
	"fmt"

	hello "github.com/ichang0301/learn-golang/1_hello-world"
	integer "github.com/ichang0301/learn-golang/2_integer"
)

func main() {
	fmt.Println(hello.Hello("world", ""))

	fmt.Println(integer.Add(1, 1))
}
