/*
 * @Author: your name
 * @Date: 2021-12-13 09:52:50
 * @LastEditTime: 2021-12-13 10:52:38
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\10-dependencyinjection.go
 */
package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// func main() {
// 	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }
