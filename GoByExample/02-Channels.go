/*
 * @Author: your name
 * @Date: 2021-12-13 09:52:50
 * @LastEditTime: 2021-12-13 22:41:22
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\GoByExample\02-Channels.go
 */
package main

import "fmt"

func main() {
	// 创建channels:                   make (chan val-type)
	messages := make(chan string)

	// send a value into a channel :  channel <-
	go func() { messages <- "ping" }()
	// receive a value from a channel:<-channel
	msg := <-messages
	// By default, sends and receives block until both the sender and receiver are ready.
	// So that we don't need any other synchronization techs.
	fmt.Println(msg)
}
