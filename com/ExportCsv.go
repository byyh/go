package com

import (
	_ "archive/zip"
	"bufio"
	"encoding/csv"
	"io"
	"logs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	_ "strconv"
	"time"
)

type ExportCsv struct {
	rows        int
	FileName    string
	ZipFileName string
	fp          *os.File
	w           *csv.Writer
	err         error
}

// 创建文件
func (this *ExportCsv) Open(fileName string) {
	this.fp, this.err = os.Create(fileName)
	CheckErr(this.err)

	this.fp.WriteString("\xEF\xBB\xBF")
	this.w = csv.NewWriter(this.fp)
	this.rows = 0
	this.FileName = fileName
}

func (this *ExportCsv) Write(data []string) {
	this.w.Write(data)
	this.rows = this.rows + 1
}

func (this *ExportCsv) Flush() {
	this.w.Flush()
}

func (this *ExportCsv) Close() {
	this.w.Flush()
	defer this.fp.Close()
}

func (this *ExportCsv) Fp() *os.File {
	return this.fp
}

func (this *ExportCsv) Remove() {
	if "" != this.FileName {
		del := os.Remove(this.FileName)
		if del != nil {
			logs.Debug("delete file failed", this.FileName)
		}
	}
	if "" != this.ZipFileName {
		del := os.Remove(this.ZipFileName)
		if del != nil {
			logs.Debug("delete file failed", this.ZipFileName)
		}
	}
}

// output  controller.Ctx.Output
func (this *ExportCsv) Download(input *http.Request, output http.ResponseWriter, filename ...string) {
	if _, err := os.Stat(file); err != nil {
		http.ServeFile(output, input, file)
		return
	}

	var fName string
	if len(filename) > 0 && filename[0] != "" {
		fName = filename[0]
	} else {
		fName = filepath.Base(file)
	}
	//https://tools.ietf.org/html/rfc6266#section-4.3
	fn := url.PathEscape(fName)
	if fName == fn {
		fn = "filename=" + fn
	} else {
		/**
		  The parameters "filename" and "filename*" differ only in that
		  "filename*" uses the encoding defined in [RFC5987], allowing the use
		  of characters not present in the ISO-8859-1 character set
		  ([ISO-8859-1]).
		*/
		fn = "filename=" + fName + "; filename*=utf-8''" + fn
	}
	output.Header("Content-Disposition", "attachment; "+fn)
	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")
	http.ServeFile(output, input, file)
}

// 压缩
func (this *ExportCsv) Compress(destFile string) {
	Zip(this.FileName, destFile)

	this.ZipFileName = destFile
}

// 输出到网络
// output  controller.Ctx.Output
func (this *ExportCsv) NetDownload(w http.ResponseWriter, filename ...string) {
	var sName string
	var fName string

	if "" != this.ZipFileName {
		sName = filepath.Base(this.ZipFileName)
	} else {
		sName = filepath.Base(this.FileName)
	}

	if len(filename) > 0 && filename[0] != "" {
		fName = filename[0]
	} else {
		fName = sName
	}

	fi, err := os.Open(sName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+url.PathEscape(fName))
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")

	r := bufio.NewReader(fi)

	buf := make([]byte, 1024)
	for i := 1; 0 < i; i = i + 1 {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}

		w.Write(buf)
		//logs.Debug(i)
		if 0 == i%500 {
			time.Sleep(time.Duration(500) * time.Millisecond)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			} else {
				logs.Debug("write flusher failed...")
				break
			}
		}
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	// output.Header("Content-Disposition", "attachment; filename="+url.PathEscape(fName))
	// output.Header("Content-Description", "File Transfer")
	// output.Header("Content-Type", "application/octet-stream")
	// output.Header("Content-Transfer-Encoding", "binary")
	// output.Header("Expires", "0")
	// output.Header("Cache-Control", "must-revalidate")
	// output.Header("Pragma", "public")

}
