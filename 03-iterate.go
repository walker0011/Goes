/*
 * @Author: your name
 * @Date: 2021-12-07 11:15:02
 * @LastEditTime: 2021-12-08 10:32:40
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\03-iterate.go
 */

package main

const repeatCount = 5

// func Repeat(a string) string {
// 	var repeated string
// 	for i := 0; i < 5; i++ {
// 		repeated = repeated + a
// 	}
// 	return repeated
// }

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
