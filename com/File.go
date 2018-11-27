package com

import (
    _"strconv"
    "os"
    "io"
    "io/ioutil"
    _"bufio"
    "github.com/astaxie/beego/logs"
)


type File struct {
    fp *os.File
    err error
}

func (this *File) Open(filename string) bool {
    _, err := os.Stat(filename)
    if nil != err {
        if os.IsNotExist(err) {
            logs.Debug("file is not exists")
            return this.Create(filename)
        }
    }

    this.fp, this.err = os.OpenFile(filename, os.O_RDWR, 0777)
    if this.err != nil {
        logs.Error("Open", this.err)
        return false
    }

    return true
}    

func (this *File) Create(filename string) bool {
    this.err = os.MkdirAll(FilePath(filename), 0777)
    if(nil != this.err) {
        logs.Error("Create dir failed:", this.err)
        return false
    }

    this.fp, this.err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
    if(nil != this.err) {
        logs.Error("Create", this.err)
        return false
    }
    
    return true
}

func (this *File) Close() {
    this.fp.Close()
}

func (this *File) WriteString(str string) int {
    n, err := this.fp.WriteString(str)
    if(nil != err) {
        logs.Error("WriteString", err)
        return 0
    }
    
    this.fp.Sync()

    return n
}   

func (this *File) WriteByte(str string) int {
    n, err := this.fp.Write([]byte(str))
    if(nil != err) {
        logs.Error(err)
        return 0
    }
    
    this.fp.Sync()

    return n
}   

func (this *File) ReadString(len int) string {
    buf := make([]byte, len)
    _, err := this.fp.Read(buf)
    if err != nil && err != io.EOF {
        logs.Debug(err)
        return ""
    }

    return string(buf)
} 

func (this *File) ReadAll(filename string) string {
    fi, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer fi.Close()

    fd, err := ioutil.ReadAll(fi)
    if nil != err {
        return ""
    }

    return string(fd)
}
