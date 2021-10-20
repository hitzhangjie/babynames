package score

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func CalcNameScore(lastName, firstName string) (float64, error) {
	api := "https://www.xingming.com/dafen/"

	p := url.Values{}
	p.Add("xs", lastName)
	p.Add("mz", firstName)
	p.Add("action", "test")

	rsp, err := http.PostForm(api, p)
	if err != nil {
		return 0, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Printf("http statusCode:%d, status:%s", rsp.StatusCode, rsp.Status)
	}

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}
	//log.Printf("http rsp body: %s", string(b))

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(b))
	if err != nil {
		return 0, err
	}

	sel := doc.Find("body > div.main.clear > div > div > p:nth-child(10) > b > font")
	//log.Printf("name's score: %s", sel.Text())

	score, err := strconv.ParseFloat(sel.Text(), 64)
	if err != nil {
		return 0, err
	}
	return score, nil
}
