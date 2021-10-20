package dup

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// CheckDuplicateNames 通过人人网检查重名人数
func CheckDuplicateNames(name string) (int, error) {
	cli := http.Client{
		Timeout: time.Second * 5,
	}

	api := "http://name.renren.com/tongMing/search"

	vals := url.Values{}
	vals.Add("q", name)
	vals.Add("cx", "014540359382904656588:9tf8clwp-ki")
	vals.Add("ie", "UTF-8")
	dat := bytes.NewBufferString(vals.Encode())

	req, err := http.NewRequest(http.MethodPost, api, dat)
	if err != nil {
		return 0, err
	}

	rsp, err := cli.Do(req)
	if err != nil {
		return 0, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("http statusCode=%d, status=%s", rsp.StatusCode, rsp.Status)
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}

	fmt.Printf("http resp body: %s\n", string(buf))

	return 0, nil
}
