/*
 * @Author: your name
 * @Date: 2021-12-06 21:12:05
 * @LastEditTime: 2021-12-07 10:31:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\01-helloworld-test.go
 */
package main

import "testing"

// test01
// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello world"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// test02
// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// test03
// func TestHello(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// 	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

// test04
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
}
