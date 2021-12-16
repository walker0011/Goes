package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func ConfigurableRace(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}

// a b 分别是URL地址
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRace(a, b, tenSecondTimeout)
}

// why struct{}: a chan struct{} is the smallest data type avaliable from a memory
// perspective so we get no allocation versus a bool.
// we are closing not sending anything on the chan
func ping(url string) chan struct{} {

	// 这里必须使用make，否则的话会zero-valued。而对于chan，zero value就是nil。
	// 对于nil 使用<-则会一直阻塞，因为无法send to nil channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch) // send a signal into the channel
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
