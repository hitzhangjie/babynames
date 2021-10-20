package radicals

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hitzhangjie/babynames/conf"
)

// ChineseRadicals 中文字符的偏旁映射表
var ChineseRadicals = map[string]string{}

// FindRadical 查找指定汉字对应的偏旁
func FindRadical(ch string) (string, bool) {

	once.Do(func() {
		m, err := loadRadical()
		if err != nil {
			panic(err)
		}
		ChineseRadicals = m
	})

	v, ok := ChineseRadicals[ch]
	return v, ok
}

var once = &sync.Once{}

func loadRadical() (map[string]string, error) {
	d := filepath.Join(conf.CfgPath, "xinhua.dict")

	fin, err := os.Open(d)
	if err != nil {
		return nil, err
	}

	var m = map[string]string{}

	sc := bufio.NewScanner(fin)
	for sc.Scan() {
		str := strings.TrimSpace(sc.Text())
		vals := strings.Split(str, ":")
		m[vals[0]] = vals[1]
	}

	return m, nil
}
