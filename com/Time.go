package com

import (
	"fmt"
	_ "math/rand"
	"os"
	"strconv"
	"time"
)

type Time struct {
	err error
}

func (mytime *Time) checkErr(param ...interface{}) {
	if mytime.err != nil {
		fmt.Println("error: ", mytime.err, param) //deal error here
		os.Exit(-1)
	}
}

// 获取昨天的时间
func (mytime *Time) Yesterday() string {

	return mytime.Add(0, 0, -1)
}

// 获取时间的时间戳
func (mytime *Time) ToTimestamp(str_time string) int64 {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", str_time, time.Local)
	ts := tm.Unix()

	return ts
}

// 获取当前本地时间
func (mytime *Time) Now() string {
	cq, _ := time.LoadLocation("Asia/Chongqing")

	return time.Now().In(cq).Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func (mytime *Time) NowTimestamp() int64 {

	return time.Now().Unix()
}

// 根据时间戳获取本地时间
func (mytime *Time) Format(timestamp int64) string {
	t := time.Unix(timestamp, 0)

	cq, _ := time.LoadLocation("Asia/Chongqing") //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。

	return t.In(cq).Format("2006-01-02 15:04:05")
}

// 根据时间差
func (mytime *Time) Add(year_diff int, month_diff int, day_diff int) string {
	cq, _ := time.LoadLocation("Asia/Chongqing") //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
	t := time.Now().AddDate(year_diff, month_diff, day_diff)

	return t.In(cq).Format("2006-01-02 03:04:05")
}

//
// 获取某个时间的时间差
// dt 原时间  2018-10-10 10:10:10
// second 相差的秒数
func (mytime *Time) Modify(dt string, second int) string {
	cq, _ := time.LoadLocation("Asia/Chongqing")                         //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
	tm, _ := time.ParseInLocation("2006-01-02 03:04:05", dt, time.Local) // time.Local

	m, _ := time.ParseDuration(strconv.Itoa(second) + "s")
	t := tm.Add(m)

	return t.In(cq).Format("2006-01-02 03:04:05")
}

//
// 获取某个时间的时间差
// dt 原时间  2018-10-10 10:10:10
// second 相差的秒数
func (mytime *Time) ModifyReturnIntDate(dt string, second int) string {
	cq, _ := time.LoadLocation("Asia/Chongqing")                         //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
	tm, _ := time.ParseInLocation("2006-01-02 03:04:05", dt, time.Local) // time.Local

	m, _ := time.ParseDuration(strconv.Itoa(second) + "s")
	t := tm.Add(m)

	return t.In(cq).Format("20060102")
}
