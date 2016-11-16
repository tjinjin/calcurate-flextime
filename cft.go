package main

import (
	"fmt"
	"time"
)

func worddayIs(date time.Time) {
	fmt.Println(date)
}

func main() {
	t := time.Now()
	startTime := time.Date(2016, 11, 01, 0, 0, 0, 0, time.Local)
	currentTime := time.Date(2016, 11, 10, 0, 0, 0, 0, time.Local)
	duration := currentTime.Sub(startTime)
	fmt.Println(startTime)
	fmt.Println(currentTime)
	fmt.Printf("%T\n", currentTime)
	fmt.Println(duration)
	fmt.Println(duration.Hours() / 24)
	workdayIs(currentTime)
	output()

}

func output() {
	fmt.Println("今月の規定flextimeはXXです")
	fmt.Println("今日時点での規定flextimeはXXです")
}
