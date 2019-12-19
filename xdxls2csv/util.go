package xdxls2csv

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

func getTime51(year int) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, 5, 1, 0, 0, 0, 0, location)
}

func getTime101(year int) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, 10, 1, 0, 0, 0, 0, location)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getClassStartTime(classNum int, isSummerTime bool) string {
	switch classNum {
	case 1:
		return "8:30"
	case 3:
		return "10:25"
	case 5:
		if isSummerTime {
			return "14:30"
		}
		return "14:00"
	case 7:
		if isSummerTime {
			return "16:25"
		}
		return "15:55"
	case 9:
		if isSummerTime {
			return "19:30"
		}
		return "19:00"
	default:
		reportFormatError()
	}
	return ""
}

func getClassEndTime(classNum int, isSummerTime bool) string {
	switch classNum {
	case 2:
		return "10:05"
	case 4:
		return "12:00"
	case 6:
		if isSummerTime {
			return "16:05"
		}
		return "15:35"
	case 8:
		if isSummerTime {
			return "18:00"
		}
		return "17:30"
	case 10:
		if isSummerTime {
			return "21:05"
		}
		return "20:35"
	default:
		reportFormatError()
	}
	return ""
}

func convWeekToI(day string) int {
	switch day {
	case "星期一":
		return 1
	case "星期二":
		return 2
	case "星期三":
		return 3
	case "星期四":
		return 4
	case "星期五":
		return 5
	default:
		reportFormatError()
	}
	return 0
}

func reportFormatError() {
	panic("课程表格式错误")
}

func exportCSV(records [][]string, out io.Writer) {
	writer := csv.NewWriter(out)
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
}
