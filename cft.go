package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	"github.com/tcnksm/go-holidayjp"
)

var now time.Time

func main() {

	vacation := flag.Float64("v", 0, "Specify number of days paid vacation. [float64]")
	flag.Parse()

	now = time.Now()
	startDay := getStartDay(now)
	//	endDay := getEndDay(now)

	// 今月のflextime計算
	var workdayCount int
	var todayCount int

	// 何日目かを取得
	duration := math.Ceil(now.Sub(startDay).Hours() / 24)
	fmt.Println(duration)
	days := make([]time.Time, int(duration))

	// 日付のsliceを作る
	for i, _ := range days {
		days[i] = startDay.AddDate(0, 0, i)
	}

	var t, w int

	for i, _ := range days {
		go countWorkday(days[i], t, w)
		todayCount += <-t
		workdayCount += <-w
	}
	output(workdayCount, todayCount, *vacation)
}

func countWorkday(tempDay time.Time, todayCount chan int, workdayCount chan int) {
	if isWorkday(tempDay) {
		if tempDay.Before(now) {
			todayCount <- 1
		}
		workdayCount <- 1
	}
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
