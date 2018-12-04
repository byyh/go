package com

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"strings"
	"errors"
	"archive/zip"

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

// 解压缩
func UnZip(zipFile string, destPath string) descNames []string {
	reader, err := zip.OpenReader(zipFile)
	CheckErr(err)
	
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		CheckErr(err)

		defer rc.Close()

		textQuoted := strconv.QuoteToASCII(file.Name)
        filename := textQuoted[1 : len(textQuoted)-1]
        filename = strings.Replace(filename, "\\", "", -1)
        filename = destPath + "/" + filename
        
		f := new(File)
		res := f.Create(filename)
		if true != res {
			fmt.Println("打开文件失败:", filename)
			panic(errors.New("打开文件失败:" + filename))
		}
		defer f.Close()
		_, err = io.Copy(f.fp, rc)
		CheckErr(err)

		f.Close()
		rc.Close()

		descNames = append(descNames, filename)
	}

	return
}

func GetDir(filepath string) string {
	return Substr(filepath, 0, strings.LastIndex(filepath, "/"))
}
