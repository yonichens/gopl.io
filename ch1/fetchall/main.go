// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)
/*
Program arguments:
https://godoc.org
https://golang.org
http://gopl.io
https://bing.com
 */

//TODO: 练习1.11：在fatchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。如果一个网站没有回应，
// 程序将采取怎样的行为？（Section8.9描述了在这种情况下的应对机制）。
func main() {
	start := time.Now()
	//ch := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan string)
	for _, url := range os.Args[1:] {
		//go fetch(url, ch) // start a goroutine
		go fetch(url, ch1) // start a goroutine
		go fetch(url, ch2) // start a goroutine
	}
	for range os.Args[1:] {
		//fmt.Println(<-ch) // receive from channel ch
		fmt.Println(<-ch1) // receive from channel ch
		fmt.Println(<-ch2) // receive from channel ch
		fmt.Println("-------------------------------")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	//只获取Body内容的字节数，其余拷贝到“垃圾桶”中
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
