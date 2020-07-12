package main

import "fmt"

type Base struct {
	name string
}

// anonymous composition = inherit
type Foo struct {
	Base
	name string
}

func main() {
	a := new(Foo)
	a.name = "www"
	a.Base.name = "eee"
	fmt.Println(a.name)
	fmt.Println(a.Base.name)
}
