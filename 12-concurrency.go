package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// 匿名函数的好处：
		//  1. 可以在声明的时候同时执行
		//  2. 所有在声明匿名函数出可用的变量在匿名函数体内都可用

		go func() {
			results[url] = wc(url)
		}()
	}
	return results
}
