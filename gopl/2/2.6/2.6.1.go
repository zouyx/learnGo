package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("begin")
	var urls = []string{"http://baidu.com", "http://163.com", "http://sina.com", "https://golang.org"}
	start := time.Now()

	ch := make(chan string)
	ch1 := make(chan string)
	for _, url := range urls[0:] {
		go fetch(url, ch)
	}

	//練習 1.10： 找一個數據量比較大的網站，用本小節中的程序調研網站的緩存策略，對每個URL執行兩遍請求，査看兩次時間是否有較大的差别，併且每次獲取到的響應內容是否一致，脩改本節中的程序，將響應結果輸出，以便於進行對比。
	for _, url := range urls[0:] {
		go fetch(url, ch1)
	}

	for range urls[0:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch1)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("end")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
