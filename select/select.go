package main

import "net/http"

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		winner = a
	case <-ping(b):
		winner = b
	}
	return
}

func ping(url string) chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		http.Get(url)
		ch <- struct{}{}
	}()
	return ch
}
