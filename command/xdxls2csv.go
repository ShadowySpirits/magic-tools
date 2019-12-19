package main

import (
	"flag"
	"fmt"
	"github.com/ShadowySpirits/magic-tools/xdxls2csv"
	"os"
	"time"
)

var startDate = flag.String("d", "", "指定本学期第一周的星期一日期，格式：2019-08-26")
var in = flag.String("i", "我的课表.xls", "指定要转换的课表路径")
var out = flag.String("o", "我的课表.csv", "指定输出路径")
var help = flag.Bool("h", false, "打印帮助")

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	flag.Usage = usage
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	check()
	xls2csv()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of xdxls2csv\n")
	flag.PrintDefaults()
}

func check() {
	if *startDate == "" {
		panic("本学期第一周的星期一日期必须指定，请使用 -help 查看帮助")
	}
}

func xls2csv() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	file, err := os.Create(*out)
	if err != nil {
		panic("创建文件失败，请检查权限")
	}
	day, err := time.ParseInLocation(xdxls2csv.TimeLayout, *startDate, location)
	if err != nil {
		panic("日期格式错误")
	}
	xdxls2csv.ParseXlsFromFileName(*in, day, file)
}
