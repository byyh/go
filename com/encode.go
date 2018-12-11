package com

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "golang.org/x/text/encoding/simplifiedchinese" 
    "golang.org/x/text/transform"
)

func GbkToUtf8(str string) (string)  {
    in := bytes.NewReader([]byte(str))
    out := transform.NewReader(in, simplifiedchinese.GBK.NewDecoder())
    data, err := ioutil.ReadAll(out)
    if err != nil {
        fmt.Println("GbkToUtf8 failed: ", err)
        return "" 
    }

    return string(data) 
}