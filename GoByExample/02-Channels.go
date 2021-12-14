/*
 * @Author: your name
 * @Date: 2021-12-13 09:52:50
 * @LastEditTime: 2021-12-13 22:41:22
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\GoByExample\02-Channels.go
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// ping function only accept a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong function accept one channel for receives(pings)
// and a second channel for sends(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// func main() {
// 	// // 创建channels:                   make (chan val-type)
// 	// messages := make(chan string)

// 	// // send a value into a channel :  channel <-
// 	// go func() { messages <- "ping" }()
// 	// // receive a value from a channel:<-channel
// 	// msg := <-messages
// 	// // By default, sends and receives block until both the sender and receiver are ready.
// 	// // So that we don't need any other synchronization techs.
// 	// fmt.Println(msg)

// 	// 2. buffered channel
// 	// messages := make(chan string, 2)

// 	// messages <- "buffered"
// 	// messages <- "channel"

// 	// fmt.Println(<-messages)
// 	// fmt.Println(<-messages)

// 	// 3. channel synchronization
// 	// done := make(chan bool , 1) // buffered channel

// 	// go worker(done)

// 	// // block until we receive a notification from the worker on the channel
// 	// <-done

// 	pings := make(chan string, 1)
// 	pongs := make(chan string, 1)

// 	ping(pings, "passed message")
// 	pong(pings, pongs)
// 	fmt.Println(<-pongs)

// }
