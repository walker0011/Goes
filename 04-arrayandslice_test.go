/*
 * @Author: your name
 * @Date: 2021-12-08 10:33:11
 * @LastEditTime: 2021-12-08 15:08:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\04-arrayandslice_test.go
 */
package main

import (
	"reflect"
	"testing"
)

// func TestSum(t *testing.T) {
// 	numbers := [...]int{1, 2, 3, 4, 5}
// 	got := Sum(numbers)
// 	want := 15

// 	if got != want {
// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
// 	}

// }

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}

func TestAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) { // not type safe
		t.Errorf("got %v want %v", got, want)
	}
}

// func TestSumAllTails(t *testing.T) {
// 	t.Run("make the sums of some slices", func(t *testing.T) {
// 		got := SumAllTails([]int{1, 2}, []int{0, 9})
// 		want := []int{2, 9}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v wnat %v", got, want)
// 		}
// 	})

// 	t.Run("safely sum empty slices", func(t *testing.T) {
// 		got := SumAllTails([]int{}, []int{0, 9})
// 		want := []int{0, 9}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v wnat %v", got, want)
// 		}
// 	})
// }

// REFACTOR

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
