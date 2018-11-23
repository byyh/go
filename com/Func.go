package com

import (
	"fmt"
	_ "math/rand"
	"strconv"
	"reflect"
)

// start从0开始
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	count := len(rs)

	if start < 0 || start > count {
		panic("start is wrong")
	}

	if length < 0 || (start+length) > count {
		panic("end is wrong")
	}

	return string(rs[start:(start + length)])
}

func CheckErr(err error, param ...interface{}) {
	if err != nil {
		fmt.Println("error: ", err, param) //deal error here
		panic(err)
	}
}

// 转换时间格式Y-m-d  Ymd
func TransDateToIntDate(date string) string {
	return Substr(date, 0, 4) + Substr(date, 5, 2) +
		Substr(date, 8, 2)
}

// 转换时间格式Ymd -> Y-m-d
func TransIntDateToDate(date string) string {
	return Substr(date, 0, 4) + "-" + Substr(date, 4, 2) +
		"-" + Substr(date, 6, 2)
}

func FloatToString(in float64) (out string) {
	out = strconv.FormatFloat(in, 'e', 2, 64)
	return
}

func Typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}
