README
------------------------------------------------------------------------------

该项目用于帮助为宝宝起名，使用方式如下：

```bash
$ git clone https://github.com/hitzhangjie/babynames
$ cd babynames
$ go install -v

$ ./babynames -h
Usage of babynames:
  -birth string
        指定出生时间 (default "2021-10-13 15:58")
  -check-dupnum
        是否检查重名
  -drop-nonused
        是否过滤掉未使用名 (default true)
  -excludes string
        排除汉字列表
  -includes string
        指定包含的汉字
  -lastname string
        指定姓氏 (default "张")
  -min-score int
        名字最低评分 (default 90)
  -name-kind string
        生成的名字类型 (default "girl_double")
  -pref-fate string
        偏好的命格 (default "木")
  -province string
        指定出生省份 (default "湖北")
  -region string
        指定出生城市 (default "武汉")
  -sex string
        指定性别, 男/女 (default "女")
```