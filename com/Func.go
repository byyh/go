package com

import (
	"fmt"
	"strings"
	_ "math/rand"
	"strconv"
	"reflect"
	"crypto/md5"
	"crypto/sha256"
    "encoding/hex"
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

// 截断文件
func FilePath(filename string) string {
    rs := []byte(filename)
    count := len(rs)

    pos := 0
    for i := count-1; i>=0; i--  {
        if("/" == string(rs[i])) {
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

func Typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}

func Md5(str string) string {
	h := md5.New()
    h.Write([]byte(str)) // 需要加密的字符串为 str
    byteMd5 := h.Sum(nil) 
    strMd5 := hex.EncodeToString(byteMd5)

    return strings.ToUpper(strMd5)
}

func Sha256Code(str string, key string) string {
    h := sha256.New()
    h.Write([]byte(str))
    byteRet := h.Sum([]byte(key))
    strReg := hex.EncodeToString(byteRet)

    return strings.ToUpper(strReg)
}