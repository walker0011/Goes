/*
 * @Author: your name
 * @Date: 2021-12-13 10:45:46
 * @LastEditTime: 2021-12-13 15:06:02
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\11-mocking.go
 */

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// spies are a kind of mock which can record how a dependency is used
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

// 用于记录sleep与fmt的顺序
type CountdownOperationSpy struct {
	Calls []string
}

// sleep一次，就添加一个“sleep”字符串
func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// write一次，就添加一个“write”字符串
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
