package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

var count = flag.Int("count", 10, "指定字库拉取次数(1次120个名字)")
var sout = flag.String("single-out", "single.out", "指定单名字库输出文件")
var dout = flag.String("double-out", "double.out", "指定双名字库输出文件")

var doubleNames = []string{}
var singleNames = []string{}

func init() {
	flag.Parse()
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSTOP)

	for i := 1; i <= *count; i++ {
		fmt.Printf("当前获取字库次数 %d (Ctrl+C停止)\n", i)

		single, double := mustFetchNames()
		singleNames = append(singleNames, single...)
		doubleNames = append(doubleNames, double...)

		singleNames = uniq(singleNames)
		doubleNames = uniq(doubleNames)

		fmt.Printf("单名字库数量 %d\n", len(singleNames))
		fmt.Printf("双名字库数量 %d\n", len(doubleNames))

		select {
		case <-ch:
			break
		default:
		}
	}

	mustWriteFile(singleNames, *sout)
	mustWriteFile(doubleNames, *dout)
}

func uniq(names []string) []string {
	set := map[string]bool{}
	for _, n := range names {
		set[n] = true
	}

	res := []string{}
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}

func mustWriteFile(names []string, out string) {
	b := &bytes.Buffer{}
	for _, n := range names {
		fmt.Fprintf(b, "%s\n", n)
	}
	err := ioutil.WriteFile(out, b.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}

func mustFetchNames() (singleNames, doubleNames []string) {

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
	//fmt.Printf("http resp headers: \n%v\n", rsp.Header)

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

	//fmt.Println("singleNames " + strings.Repeat("-", 120))
	//fmt.Println(singleNames)
	//fmt.Println()
	//
	//fmt.Println("doubleNames " + strings.Repeat("-", 120))
	//fmt.Println(doubleNames)
	//fmt.Println()

	return
}
