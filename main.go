package main

import (
	"fmt"
	"log"
	"strings"
	"sync/atomic"

	"github.com/hitzhangjie/babynames/names"
	"github.com/hitzhangjie/babynames/radicals"
	"github.com/hitzhangjie/babynames/score"
	"github.com/hitzhangjie/babynames/wuxing"
	"go.uber.org/ratelimit"
)

func main() {
	// init the app
	log.SetPrefix("")
	log.SetFlags(log.Lmsgprefix)

	opts := parse()

	log.Printf("您正在使用babynames进行名字测算")
	//log.Printf("- 是否检查评分: %v", opts.score)
	log.Printf("- 是否检查重名: %v", opts.dup)
	log.Printf("- 过滤未使用名: %v", opts.drop)
	log.Printf("- 您偏好的五行: %s", opts.fates)
	log.Println()

	// generate all names
	log.Println("姓名字典生成中，请稍后...")

	names, err := generateFirstNames(opts)
	if err != nil {
		panic(err)
	}

	// 丢掉偏旁不符合偏好五行要求的
	names, err = filterByRadicals(opts.fates, names)
	if err != nil {
		panic(err)
	}

	// 丢掉不包含必选字的
	names, err = filterByIncludes(opts.includes, names)
	if err != nil {
		panic(err)
	}

	// 丢掉要排除掉字的
	names, err = filterByExcludes(opts.excludes, names)
	if err != nil {
		panic(err)
	}
	//log.Printf("%s\n", beautify(names))

	// 计算评分
	names, err = filterByScore(opts, names, float64(opts.minScore))
	if err != nil {
		panic(err)
	}
	log.Printf("\n%s\n", beautify(names))
}

// 根据NameKind获取单字、双字候选名字列表
func generateFirstNames(opt *options) ([]string, error) {
	return names.FindNameDict(names.NameKind(opt.kind))
}

// 过滤掉候选名字偏旁不符合五行要求的
func filterByRadicals(fates []string, names []string) ([]string, error) {
	var candicates []string

	for _, n := range names {
		for _, fate := range fates {
			if !CheckRadicals(fate, n) {
				continue
			}
			candicates = append(candicates, n)
			break
		}
	}

	return candicates, nil
}

// 过滤掉不包含必选字的
func filterByIncludes(chs []string, names []string) ([]string, error) {
	if len(chs) == 0 {
		return names, nil
	}

	var candicates []string
	for _, n := range names {
		included := false
		for _, ch := range chs {
			if !strings.Contains(n, ch) {
				continue
			}
			included = true
			break
		}
		if !included {
			continue
		}
		candicates = append(candicates, n)
	}
	return candicates, nil
}

// 过滤掉不包含必选字的
func filterByExcludes(chs []string, names []string) ([]string, error) {
	if len(chs) == 0 {
		return names, nil
	}

	var candicates []string
	for _, n := range names {
		dropped := false
		for _, ch := range chs {
			if !strings.Contains(n, ch) {
				continue
			}
			dropped = true
			break
		}
		if !dropped {
			candicates = append(candicates, n)
		}
	}
	return candicates, nil
}

func filterByScore(opts *options, names []string, threshold float64) ([]string, error) {
	var candicates []string

	ch := doFilterByScore(opts, names, threshold)
	for v := range ch {
		candicates = append(candicates, v)
	}
	return candicates, nil
}

func doFilterByScore(opts *options, names []string, threshold float64) chan string {

	fmt.Printf("共有 %d 个名字待计算分数\n", len(names))
	count := int64(0)

	rspCh := make(chan string, len(names))

	newTask := func(lastName, firstName string, threshold float64) func() {
		return func() {
			defer func() {
				v := atomic.AddInt64(&count, 1)
				if v > 1 {
					fmt.Printf(strings.Repeat("\b", 4)+"%4d", v)
				} else {
					fmt.Printf("当前已计算个数 %4d", v)
				}

				if int(v) == len(names) {
					close(rspCh)
				}
			}()
			ok, err := isScoreOK(lastName, firstName, threshold)
			if err != nil {
				log.Printf("isScoreOK, err: %v", err)
				return
			}
			if !ok {
				return
			}
			rspCh <- firstName
		}
	}

	limit := ratelimit.New(10)

	for _, n := range names {
		go func(lastName, firstName string) {
			limit.Take()
			newTask(lastName, firstName, threshold)()
		}(opts.lastName, n)
	}

	return rspCh
}

func isScoreOK(lastName, firstName string, threshold float64) (bool, error) {
	score, err := score.CalcNameScore(lastName, firstName)
	if err != nil {
		return false, err
	}

	return score >= threshold, nil
}

// CheckRadicals 检查firstName中是否有字的偏旁满足命格候选字要求
func CheckRadicals(fatePref, firstName string) bool {
	chs, ok := wuxing.FateChars[wuxing.FateKind(fatePref)]
	if !ok {
		return false
	}

	found := false

NextChar:
	for _, c := range firstName {
		ch := fmt.Sprintf("%c", c)
		r, ok := radicals.FindRadical(ch)
		if !ok {
			continue
		}

		for _, ch := range chs {
			if ch == r {
				found = true
				break NextChar
			}
		}
	}
	return found
}

func beautify(names []string) string {
	sb := &strings.Builder{}

	for i, n := range names {
		fmt.Fprintf(sb, "%s\t", n)
		if i > 0 && (i+1)%15 == 0 {
			fmt.Fprintf(sb, "\n")
		}
	}

	return sb.String()
}
