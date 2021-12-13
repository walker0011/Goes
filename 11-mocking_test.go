/*
 * @Author: your name
 * @Date: 2021-12-13 10:45:52
 * @LastEditTime: 2021-12-13 14:55:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\11-mocking_test.go
 */
package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spysleeper := &SpySleeper{}

	Countdown(buffer, spysleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spysleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spysleeper.Calls)
	}
}
