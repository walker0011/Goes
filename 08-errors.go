/*
 * @Author: your name
 * @Date: 2021-12-10 20:34:24
 * @LastEditTime: 2021-12-10 21:32:49
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\08-errors.go
 */

package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// func main() {
// 	for _, i := range []int{7, 42} {
// 		if r, e := f1(i); e != nil {
// 			fmt.Println("f1 failed", e)
// 		} else {
// 			fmt.Println("f1 worked", r)
// 		}
// 	}

// 	for _, i := range []int{7, 42} {
// 		if r, e := f2(i); e != nil {
// 			fmt.Println("f2 failed", e)
// 		} else {
// 			fmt.Println("f2 worked", r)
// 		}
// 	}

// 	_, e := f2(42)
// 	if ae, ok := e.(*argError); ok {
// 		fmt.Println(ae.arg)
// 		fmt.Println(ae.prob)
// 	}

// }
