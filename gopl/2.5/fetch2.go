package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//練習 1.8： 脩改fetch這個范例，如果輸入的url參數沒有 http:// 前綴的話，爲這個url加上該前綴。你可能會用到strings.HasPrefix這個函數。
	url := "baidu.com"
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	//練習 1.7： 函數調用io.Copy(dst, src)會從src中讀取內容，併將讀到的結果寫入到dst中，使用這個函數替代掉例子中的ioutil.ReadAll來拷貝響應結構體到os.Stdout，避免申請一個緩衝區（例子中的b）來存儲。記得處理io.Copy返迴結果中的錯誤。
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	//練習 1.9： 脩改fetch打印出HTTP協議的狀態碼，可以從resp.Status變量得到該狀態碼。
	fmt.Println("status code:" + resp.Status)
}
