/*
 * @Author: your name
 * @Date: 2021-12-06 20:22:13
 * @LastEditTime: 2021-12-07 10:39:33
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \projects\01-helloworld.go
 */

package main

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix

	}

	return prefix + name
	// 将return "Hello, "+ name修改为常量+name可以提升性能，因为避免了每次调用Hello函数创建string 实例的过程
}

func greetingPrefix(language string) (prefix string) { // named return value
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// // test01
// func Hello() string { // return a string
// 	return "Hello World"
// }

// func main() {
// 	fmt.Println(Hello())
// 	fmt.Println("hello world")
// }
