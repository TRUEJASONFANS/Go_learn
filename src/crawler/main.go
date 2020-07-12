package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://c0.3.cn/stock?skuId=37892094529&cat=1713,3281,3683&area=1_2803_2829_0"
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	jsonStr := Decode(string(p))
	fmt.Println(jsonStr)

	var jingweiI interface{}
	err = json.Unmarshal([]byte(jsonStr), &jingweiI)
	// 获取某个 key 的值
	jingweiM, ok := jingweiI.(map[string]interface{})
	if !ok {
		fmt.Println("DO SOMETHING!")
		return
	}
	fmt.Println(jingweiM["stock"].(map[string]interface{})["jdPrice"].(map[string]interface{})["op"])
	//searchJDWebpage("9787115238870")
}
func searchJDWebpage(ISPN string) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		startsWith := strings.HasPrefix(url, "//item.jd.com") // true
		if startsWith {
			e.Request.Visit(url)
		}
	})

	c.OnHTML(".p-discount", func(e *colly.HTMLElement) {
		value := e.Text
		if value != "" {
			fmt.Println("price:" + value)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://search.jd.com/Search?keyword=" + ISPN + "&enc=utf-8&wq=" + ISPN)
}

func Encode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
func Decode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
