package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		winner = a
	case <-ping(b):
		winner = b
	case <-time.After(timeout):
		err = errors.New(fmt.Sprintf("timed out waiting for %s and %s", a, b))
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
