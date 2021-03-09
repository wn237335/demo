package main

import (
	"time"
)

func main() {
	//fmt.Println("2021-03-01 15:46:15 +0800 CST".)
	//fmt.Println(getHourDiffer("2020-09-10 13:00:00", "2020-09-10 14:50:00"))
}

//获取相差时间
func getHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 60
		return hour
	} else {
		return hour
	}
}
