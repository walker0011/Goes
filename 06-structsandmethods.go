/*
 * @Author: your name
 * @Date: 2021-12-08 19:51:32
 * @LastEditTime: 2021-12-09 19:20:57
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\06-structs.go
 */

package main

type person struct {
	name string
	age  int
}

// TODO: you can safely return a pointer to local variable as a local variable will survive the scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// method defined on struct types

// type rect struct {
// 	width, height int
// }

// // this area method has a receiver type of rect
// func (r rect) area() int {
// 	return r.width * r.height
// }

// // this perim method has a receiver type of *rect
// func (r *rect) perim() int {
// 	return 2*r.width + 2*r.height
// }

// func main() {
// 	fmt.Println(person{"Bob", 20})

// 	fmt.Println(person{name: "Alice", age: 30})

// 	fmt.Println(person{name: "Fred"}) // omitted fields will be zero-valued

// 	fmt.Println(&person{name: "Ann", age: 40})

// 	fmt.Println(newPerson("Jon"))

// 	s := person{name: "Sean", age: 50}
// 	fmt.Println(s.name)

// 	sp := &s
// 	fmt.Println(sp.age) // pointers are automatically dereferenced

// 	sp.age = 51
// 	fmt.Println(sp.age)

// 	// Test method defined on struct
// 	r := rect{width: 10, height: 5}
// 	fmt.Println("area: ", r.area())
// 	fmt.Println("perim: ", r.perim())

// 	// Go automatically handles conversion between values and pointers for method calls.

// 	rp := &r
// 	fmt.Println("area: ", rp.area())
// 	fmt.Println("perim: ", rp.perim())

// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Height * r.Width
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius
// }

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Height + rectangle.Width)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Height * rectangle.Width
// }
