package main

import "fmt"

func main() {
	var j int = 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i, j:%d, %d\n", i, j)
		}
	}()
	a()
	fmt.Print(a)
}
