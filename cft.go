package main

import (
	"fmt"
	"math"
	"time"
)

func worddayIs(date time.Time) {
	fmt.Println(date)
}

func main() {
	startTime := time.Date(2016, 11, 01, 0, 0, 0, 0, time.Local)
	currentTime := time.Now()
	duration := currentTime.Sub(startTime)
	fmt.Println(startTime)
	fmt.Println(currentTime)
	fmt.Printf("%T\n", currentTime)
	fmt.Println(duration)
	fmt.Println(math.Ceil(duration.Hours() / 24))
	output(int(math.Ceil(duration.Hours()/24)) * 8)
}

func output(i int) {
	fmt.Printf("今月の規定flextimeは%d 時間です\n", 1000)
	fmt.Printf("今日時点の規定flextimeは%d 時間です\n", i)
}
