/*
 * @Author: your name
 * @Date: 2021-12-13 09:52:50
 * @LastEditTime: 2021-12-13 10:37:08
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\09-embedding.go
 */
package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// a container embeds a base
type container struct {
	base
	str string
}

// func main() {
// 	co := container{
// 		base: base{
// 			num: 1,
// 		},
// 		str: "some name",
// 	}

// 	fmt.Printf("co = {num:%v, str:%v}\n", co.num, co.str)

// 	fmt.Println("also num:", co.base.num)

// 	fmt.Println("describe:", co.describe())

// 	type describer interface {
// 		describe() string
// 	}

// 	var d describer = co
// 	fmt.Println("describer:", d.describe())
// }
