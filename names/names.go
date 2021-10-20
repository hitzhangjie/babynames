package names

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/hitzhangjie/babynames/conf"
)

// NameKind 取名类型
type NameKind string

const (
	GirlSingleName NameKind = "girl_single" // 女孩单字名
	GirlDoubleName NameKind = "girl_double" // 女孩双子名
	BoySingleName  NameKind = "boy_single"  // 男孩单字名
	BoyDoubleName  NameKind = "boy_double"  // 男孩双子名
)

// namesDict 名字字典
var namesDict map[NameKind][]string

func init() {
	namesDict = map[NameKind][]string{
		GirlSingleName: mustLoadNames("girl_single.dict"),
		GirlDoubleName: mustLoadNames("girl_double.dict"),
		BoySingleName:  mustLoadNames("boy_single.dict"),
		BoyDoubleName:  mustLoadNames("boy_double.dict"),
	}
}

func mustLoadNames(fp string) []string {
	d := filepath.Join(conf.CfgPath, fp)

	fin, err := os.Open(d)
	if err != nil {
		panic(err)
	}

	var names []string

	sc := bufio.NewScanner(fin)
	for sc.Scan() {
		names = append(names, strings.TrimSpace(sc.Text()))
	}

	return names
}

// FindNameDict 查找特定名字类型对应的名字字典
func FindNameDict(kind NameKind) ([]string, error) {
	v, ok := namesDict[kind]
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}
