package swmfunc

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SwmGetHtml() {

	url := "https://www.xbiqugew.com/book/55830/40751931.html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Cookie", "fontsize=12px; Hm_lvt_252f9a986eb5a291cc4f56bcecd88721=1715394738; Hm_lvt_a71b1bc761fe3f26085e79b5fd6a7f71=1715394739; Hm_lpvt_252f9a986eb5a291cc4f56bcecd88721=1715409395; Hm_lpvt_a71b1bc761fe3f26085e79b5fd6a7f71=1715409395")
	req.Header.Add("Priority", "u=0, i")
	req.Header.Add("Referer", "https://www.xbiqugew.com/book/55830/40751930.html")
	req.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
