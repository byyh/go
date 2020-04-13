package com

import (
	"encoding/json"
	"fmt"
	_ "math/rand"
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

// start从0开始
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	count := len(rs)

	if start < 0 || start > count {
		panic("start is wrong")
	}

	if length < 0 {
		panic("end is wrong")
	}

	if (start + length) > count {
		length = count - start
	}

	return string(rs[start:(start + length)])
}

// 截断文件
func FilePath(filename string) string {
	rs := []byte(filename)
	count := len(rs)

	pos := 0
	for i := count - 1; i >= 0; i-- {
		if "/" == string(rs[i]) {
			pos = i
			break
		}
	}

	return string(rs[0:pos])
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

func StringToFloat64(s string) (out float64) {
	out, _ = strconv.ParseFloat(s, 64)
	return
}

func StringToFloat32(s string) float32 {
	out, _ := strconv.ParseFloat(s, 32)
	return float32(out)
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func ClassFieldToSortUrlParameter(refT reflect.Type, refV reflect.Value, filter []string) string {
	elet := refT.Elem()
	elev := refV.Elem()
	num := elet.NumField()
	keys := []string{}
	maps := make(map[string]string)

	i := 0
	for i < num {
		field := elet.Field(i)
		value := elev.Field(i)
		i++
		//fmt.Println(field.Name, filter)
		if InArrayString(filter, field.Name) {
			continue
		}

		keys = append(keys, field.Name)
		maps[field.Name] = url.QueryEscape(value.String())
	}

	sort.Strings(keys)

	str := ""
	for _, k := range keys {
		if "" == str {
			str = k + "=" + maps[k]
		} else {
			str += "&" + k + "=" + maps[k]
		}
	}

	return str
}

func InArrayString(arr []string, value string) bool {
	for _, k := range arr {
		if k == value {
			return true
		}
	}

	return false
}

// json decode
func JsonDecode(data string) (map[string]interface{}, error) {
	var tmp interface{}
	result := make(map[string]interface{})

	err := json.Unmarshal([]byte(data), &tmp) // strings.NewReader(data)
	if nil != err {
		return result, err
	}

	fmt.Println(tmp)
	maps := tmp.(map[string]interface{})
	for k, v := range maps {
		switch Typeof(v) {
		case "float64":
			result[k] = FloatToString(v.(float64))
		case "int":
			result[k] = strconv.Itoa(v.(int))
		case "bool":
			result[k] = v.(bool)
		case "string":
			result[k] = v.(string)
		default:
			result[k] = v
		}
	}

	return result, nil
}

// 字符串数组列表排序
// arr 需要排序的字符串列表，该参数传入的是指针，排序后的值会保存在该参数中
// isForward 排序的方向，true正向从小到大；false反向从大到小
func SortStringList(arr []string, isForward bool) {
	length := len(arr)
	var tmp string
	var tag bool
	for j := 0; j < length; j++ {
		tag = false
		for k := 0; k < length-j-1; k++ {
			if isForward {
				if arr[k] > arr[k+1] {
					tmp = arr[k]
					arr[k] = arr[k+1]
					arr[k+1] = tmp
					tag = true
				}
			} else {
				if arr[k] < arr[k+1] {
					tmp = arr[k]
					arr[k] = arr[k+1]
					arr[k+1] = tmp
					tag = true
				}
			}
		}

		if !tag {
			break
		}
	}
}
