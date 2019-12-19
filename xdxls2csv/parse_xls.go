package xdxls2csv

import (
	"github.com/extrame/xls"
	"io"
	"regexp"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02"

func ParseXlsFromFileName(filename string, startDate time.Time, out io.Writer) {
	workBook, closer, err := xls.OpenWithCloser(filename, "utf-8")
	if err != nil {
		panic("课表路径错误")
	}
	defer func() {
		if err := closer.Close(); err != nil {
			panic(err)
		}
	}()
	parseXls(workBook, startDate, out)
}

func ParseXlsFromFile(file io.ReadSeeker, startDate time.Time, out io.Writer) {
	workBook, err := xls.OpenReader(file, "utf-8")
	if err != nil {
		panic("课表读取错误")
	}
	parseXls(workBook, startDate, out)
}

func parseXls(file *xls.WorkBook, startDate time.Time, out io.Writer) {
	if sheet := file.GetSheet(0); sheet != nil {
		records := [][]string{
			{"Subject", "StartDate", "EndDate", "StartTime", "EndTime", "Location", "Description"},
		}
		for i := 1; i <= int(sheet.MaxRow); i++ {
			var data []string
			row := sheet.Row(i)

			name := row.Col(1)
			day := row.Col(6)
			startNum := row.Col(7)
			endNum := row.Col(8)
			teacherName := row.Col(9)
			classNum := row.Col(0)
			location := row.Col(10)
			weeks := row.Col(5)

			duration := strings.Split(weeks, ",")
			singleOrDouble := 0
			if strings.Contains(weeks, "单") {
				singleOrDouble = 1
			} else if strings.Contains(weeks, "双") {
				singleOrDouble = 2
			}

			for _, s := range duration {
				reg := regexp.MustCompile(`\d+`)
				matches := reg.FindAllString(s, -1)
				weekStart := atoi(matches[0]) - 1
				var maxWeek int
				if len(matches) == 1 {
					maxWeek = weekStart
				} else {
					maxWeek = atoi(matches[1])
				}

				for i := weekStart; i < maxWeek; i++ {
					if singleOrDouble == 1 && i%2 != 0 {
						continue
					}
					if singleOrDouble == 2 && i%2 == 0 {
						continue
					}

					day := convWeekToI(day)
					date := startDate.AddDate(0, 0, 7*i+day-1)

					isSummerTime := date.After(getTime51(date.Year())) && date.Before(getTime101(date.Year()))
					data = append(data, name)
					data = append(data, date.Format(TimeLayout))
					data = append(data, date.Format(TimeLayout))
					data = append(data, getClassStartTime(atoi(startNum), isSummerTime))
					data = append(data, getClassEndTime(atoi(endNum), isSummerTime))
					data = append(data, location)
					data = append(data, teacherName+" "+classNum)

					records = append(records, data)
					data = []string{}
				}
			}
		}

		exportCSV(records, out)
	} else {
		reportFormatError()
	}
}
