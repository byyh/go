package com 

import ( 
    _"bytes" 
    "fmt"
    "strings"
    "errors"
    "strconv"
    "io/ioutil" 
    "net/http" 
    //"net/url" 
    )

type Http struct {
    Client *http.Client
}

func (this *Http) Post(
    domain string, 
    data string, 
    headerMaps map[string]string) (string, error) {  
    if "" == domain {
        return "", errors.New("failed: post domain is not empty")
    }
    if "" == data {
        return "", errors.New("failed: post data is not empty")
    }

    body := strings.NewReader(data)

    this.Client = &http.Client{}
  
    req, _ := http.NewRequest("POST", domain, body)

    for k, v := range headerMaps {
        //fmt.Println(k, v)
        req.Header.Set(k, v)
    }
    //req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
    //fmt.Println(req)
    resp, err := this.Client.Do(req)  // 发送    
    if err != nil { 
        fmt.Println("failed: 发送失败", err.Error()) 

        return "", err
    }    

    defer resp.Body.Close()      // 关闭resp.Body

    if 200 == resp.StatusCode {
        body, _ := ioutil.ReadAll(resp.Body) 

        return string(body), nil
    } else {
        body, _ := ioutil.ReadAll(resp.Body)

        return string(body), errors.New("failed: network error,ret status code=" + strconv.Itoa(resp.StatusCode))
    }    
}