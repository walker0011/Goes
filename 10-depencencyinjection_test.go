/*
 * @Author: your name
 * @Date: 2021-12-13 09:52:50
 * @LastEditTime: 2021-12-13 10:28:14
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\10-depencencyinjection_test.go
 */
package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // bytes.Buffer implements the Writer interface， it has the method Write
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
