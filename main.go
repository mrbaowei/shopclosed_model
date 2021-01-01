package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	OPEN   int32 = 1
	CLOSED int32 = 0
)

var Customer int64 = 0
var Status int32 = 1

func Closed() {
	if Status == OPEN {
		atomic.StoreInt32(&Status, CLOSED)
	}
	tk := time.NewTicker(time.Microsecond * 10)
	defer tk.Stop()
	for range tk.C {
		if Customer == 0 {
			break
		}
	}
	fmt.Println("shop close")
}

func Opened() {
	fmt.Println("shop open")
	if Status == CLOSED {
		atomic.StoreInt32(&Status, OPEN)
	}
}

func Welcome() {
	for {
		if Status == OPEN {
			break
		}
	}
	atomic.AddInt64(&Customer, 1)
}

func Goodbye() {
	if Customer <= 0 {
		return
	}
	atomic.AddInt64(&Customer, -1)
}

func main() {
	go func() {
		for {
			go func() {
				Welcome()
				fmt.Println("customer shopping")
				time.Sleep(1 * time.Second)
				Goodbye()
			}()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			Closed()
			fmt.Println("shop prepare")
			time.Sleep(2 * time.Second)
			Opened()
		}
	}()

	select {}
}
