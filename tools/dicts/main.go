package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

var doubleNames = []string{}
var singleNames = []string{}

func main() {
	api := "https://www.sheup.net/mingzi_girl_1.php"

	p := url.Values{}
	p.Add("xing_key", "")
	p.Add("name_key", "")
	buf := bytes.NewBufferString(p.Encode())

	req, err := http.NewRequest(http.MethodPost, api, buf)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("http statusCode:%d status:%s", rsp.StatusCode, rsp.Status)
		return
	}

	dat, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http resp headers: \n%v\n", rsp.Header)

	out, err := iconv.ConvertString(string(dat), "GB2312", "UTF-8")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("http resp body:\n%s\n", out)

	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(out))
	if err != nil {
		panic(err)
	}
	sel := doc.Find("body > div > div.s_content > div.content_main > div.main_mingzi > div > ul")
	fmt.Println("nodes: ", len(sel.Children().Nodes))
	//html, err := sel.Html()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("nodes: ", html)

	sel.Children().Each(func(idx int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		switch len([]rune(name)) {
		case 1:
			singleNames = append(singleNames, name)
		case 2:
			doubleNames = append(doubleNames, name)
		default:
		}
	})

	fmt.Println("singleNames " + strings.Repeat("-", 120))
	fmt.Println(singleNames)
	fmt.Println()

	fmt.Println("doubleNames " + strings.Repeat("-", 120))
	fmt.Println(doubleNames)
	fmt.Println()
}
