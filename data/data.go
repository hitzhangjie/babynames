package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fp := filepath.Join(dir, "诗经解读.txt")
	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	var count int

	var start bool
	var title string

	var buf = &bytes.Buffer{}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s := strings.TrimSpace(sc.Text())
		if len(s) == 0 {
			continue
		}

		if s == "【原文】" {
			start = true
			count++
		}

		if !start {
			continue
		}

		// read title
		for sc.Scan() {
			s := strings.TrimSpace(sc.Text())
			if len(s) == 0 || strings.Contains(s, "《") {
				continue
			}
			title = s
			break
		}

		fmt.Fprintf(buf, strings.Repeat("-", 78)+"\n")
		fmt.Fprintf(buf, "诗经-%d : %s\n", count, title)

		var skipNextLine = false

		// read body
		for sc.Scan() {
			s := strings.TrimSpace(sc.Text())
			if len(s) == 0 {
				continue
			}

			// 跳过注释节，保留经典原意节，读取到当代阐释节结束
			switch s {
			case "【注释】":
				skipNextLine = true
			case "【经典原意】":
				skipNextLine = false
			}
			if skipNextLine {
				continue
			}

			if s == "【当代阐释】" {
				start = false
				break
			}

			fmt.Fprintf(buf, "%s\n", s)
		}
	}

	dst := filepath.Join(dir, "诗经解读.ext.txt")
	err = ioutil.WriteFile(dst, buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
