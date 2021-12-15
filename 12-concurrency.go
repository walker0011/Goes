package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		// 匿名函数的好处：
		//  1. 可以在声明的时候同时执行
		//  2. 所有在声明匿名函数出可用的变量在匿名函数体内都可用

		// go func() {
		// 	// 注意这种方式：每个goroutine都是url的reference，因此最后写入的是最后一个url信息
		// 	results[url] = wc(url)
		// }()

		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
