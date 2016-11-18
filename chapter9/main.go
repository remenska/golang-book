package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return 2*distance(r.x1, r.y1, r.x1, r.y2) + 2*distance(r.x1, r.y1, r.x2, r.y1)
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

type Android struct {
	Person
	model string
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultyShape struct {
	shapes []Shape
}

func (m *MultyShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	// var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	// fmt.Println(circleArea(cx, cy, cr))
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	a := new(Android)
	a.name = "Daniela"
	a.talk()
	a.Person.name = "Hodor"
	a.talk()
	fmt.Println(totalArea(&c, &r))
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}
