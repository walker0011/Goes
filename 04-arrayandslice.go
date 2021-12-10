/*
 * @Author: your name
 * @Date: 2021-12-08 10:33:20
 * @LastEditTime: 2021-12-08 15:06:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\04-arrayandslice.go
 */
package main

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sum := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
