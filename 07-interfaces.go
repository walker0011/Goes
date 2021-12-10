/*
 * @Author: your name
 * @Date: 2021-12-09 19:09:11
 * @LastEditTime: 2021-12-09 21:23:26
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\04-interfaces.go
 */
package main

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface {
	Area() float64
}

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Height + rectangle.Width)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Height * rectangle.Width
// }

// type geometry interface {
// 	area() float64
// 	perim() float64
// }

// type rect struct {
// 	width, height float64
// }

// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (r rect) perim() float64 {
// 	return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// func main() {
// 	r := rect{width: 3, height: 4}
// 	c := circle{radius: 5}

// 	measure(r)
// 	measure(c)
// }
