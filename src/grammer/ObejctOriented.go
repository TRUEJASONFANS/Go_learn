package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}
type Circle struct {
	r float64
}

func (r *Rect) Area() float64 {
	fmt.Println(r)
	r.width = 20.0
	return r.width * r.height
}
func (r Rect) Area2() float64 {
	fmt.Println(r)
	r.width = 30
	return r.width * r.height
}
func NewRect(x, y, width, height float64) *Rect {
	//return &Rect{x, y, width, height}
	return new(Rect)
}

func (c *Circle) area() float64 {
	return c.r * c.r
}

func main() {
	a := NewRect(1, 2, 10, 20)
	area := a.Area()
	area2 := a.Area2()
	b := &Circle{r: 30}

	fmt.Println("Area: %f, width: %f", area, a.width)
	fmt.Println("Area2: %f, Area2: %f", area2, a.width)
	fmt.Println("Cirle's area", b.area())

	var v1 interface{} = a
	a2, ok := v1.(*Rect)
	if ok {
		fmt.Printf("I am a2's area :%f", a2.Area())
	}

}
