/*
 * @Author: your name
 * @Date: 2021-12-08 19:47:13
 * @LastEditTime: 2021-12-08 19:57:02
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\05-pointers.go
 */

// from : https://gobyexample.com/pointers
package main

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// func main() {
// 	i := 1
// 	fmt.Println("initial:", i)

// 	zeroval(i)
// 	fmt.Println("zeroval:", i)

// 	zeroptr((&i))
// 	fmt.Println("zeroptr:", i)

// 	fmt.Println("pointer:", &i) // output the address of local variable i

// }
