package com

import (
	_"strconv"
	"os"
	"io"
	_"bufio"
	"archive/zip"
	_"github.com/astaxie/beego/logs"
)

// 压缩
func Zip(sourceFile string, destFile string) {
	s, err := os.Open(sourceFile)
	CheckErr(err)
	d, err := os.Create(destFile)
	CheckErr(err)

	defer s.Close()
	defer d.Close()

	zw := zip.NewWriter(d)
	defer zw.Close()

	info, err := d.Stat()
	CheckErr(err)

	header, err := zip.FileInfoHeader(info)	
	CheckErr(err)
	
	header.Name = sourceFile
	header.Method = zip.Deflate
	writer, err := zw.CreateHeader(header)
	CheckErr(err)

	_, err = io.Copy(writer, s)
	CheckErr(err)
}