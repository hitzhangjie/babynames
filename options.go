package main

import (
	"flag"
	"strings"
	"time"
)

var (
	kind = flag.String("name-kind", "girl_double", "生成的名字类型")
	fate = flag.String("pref-fate", "木", "偏好的命格")

	minScore = flag.Int("min-score", 90, "名字最低评分")
	dup      = flag.Bool("check-dupnum", false, "是否检查重名")
	drop     = flag.Bool("drop-nonused", true, "是否过滤掉未使用名")

	lastName = flag.String("lastname", "张", "指定姓氏")
	includes = flag.String("includes", "", "指定包含的汉字")
	excludes = flag.String("excludes", "", "排除汉字列表")
	sex      = flag.String("sex", "女", "指定性别, 男/女") //deprecated
	birth    = flag.String("birth", "2021-10-13 15:58", "指定出生时间")

	province = flag.String("province", "湖北", "指定出生省份")
	region   = flag.String("region", "武汉", "指定出生城市")
)

type options struct {
	minScore int
	fates    []string
	dup      bool
	drop     bool
	kind     string

	lastName string
	includes []string
	excludes []string
	sex      string
	birth    time.Time

	province string
	region   string
}

const datetime = "2006-01-02 15:04"

func parse() *options {
	flag.Parse()

	var options options

	options.fates = strings.Split(*fate, "")
	options.kind = *kind

	options.minScore = *minScore
	options.dup = *dup
	options.drop = *drop

	options.lastName = *lastName
	options.includes = strings.Split(*includes, "")
	options.excludes = strings.Split(*excludes, "")
	options.sex = *sex

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	t, err := time.ParseInLocation(datetime, *birth, loc)
	if err != nil {
		panic(err)
	}
	options.birth = t

	options.province = *province
	options.region = *region

	return &options
}
