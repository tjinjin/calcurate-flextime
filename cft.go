package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tcnksm/go-holidayjp"
)

var now time.Time

func main() {

	salaried := flag.Float64("s", 0, "Specify number of days paid vacation. [float64]")
	flag.Parse()

	now = time.Now()
	startDay := getStartDay(now)
	endDay := getEndDay(now)

	// 今月のflextime計算
	var workdayCount int
	var todayCount int
	var tempDay time.Time

	tempDay = startDay

	for {
		if endDay.Before(tempDay) {
			break
		}
		if isWorkday(tempDay) {
			if tempDay.Before(now) {
				todayCount += 1
			}
			workdayCount += 1
		}
		//１日進める
		tempDay = tempDay.Add(24 * time.Hour)
	}
	output(workdayCount, todayCount, *salaried)
}

// i 日数
func output(i, j int, k float64) {
	fmt.Printf("今月の規定出勤日数は %d 日です\n", i)
	fmt.Printf("今月の規定労働時間は %d 時間です\n", i*8-int(k*8))
	fmt.Printf("%d日時点の標準労働時間は %d 時間です\n", now.Day(), j*8-int(k*8))
}

// 21 - 20の周期
// 11/20 -> 10/21
// 11/21 -> 11/21
//  1/20 -> 12/21
func getStartDay(d time.Time) time.Time {
	var startDay time.Time
	switch n := d.Day(); {
	case (n >= 21) && (n <= 31):
		startDay := time.Date(d.Year(), d.Month(), 21, 0, 0, 0, 0, time.Local)
		return startDay
	case (n >= 1) && (n <= 20):
		// 前月の情報を取りたい
		tmp := d.AddDate(0, 0, -21)
		startDay := time.Date(tmp.Year(), tmp.Month(), 21, 0, 0, 0, 0, time.Local)
		return startDay
	}
	return startDay
}

// 11/21 -> 12/20
// 11/1  -> 11/20
// 12/21 ->  1/20
func getEndDay(d time.Time) time.Time {
	var endDay time.Time
	switch n := d.Day(); {
	case (n >= 21) && (n <= 31):
		// 翌月の情報を取りたい
		tmp := d.AddDate(0, 0, 11)
		endDay := time.Date(tmp.Year(), tmp.Month(), 20, 0, 0, 0, 0, time.Local)
		return endDay
	case (n >= 1) && (n <= 20):
		endDay := time.Date(d.Year(), d.Month(), 20, 0, 0, 0, 0, time.Local)
		return endDay
	}
	return endDay
}

func isWorkday(d time.Time) bool {
	workday := true
	//土日チェック
	// Weedkay 0 Sun 1 Mon ... 6 Sat
	if (d.Weekday() == 0) || (d.Weekday() == 6) {
		workday = false
	}

	//Holidayチェック
	if holidayjp.IsHoliday(d) {
		workday = false
	}
	return workday
}
