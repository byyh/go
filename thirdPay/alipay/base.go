package alipay 

import ( 
    _"bytes" 
    "fmt"
    "errors"
    "io/ioutil" 
    "net/http" 
    "net/url" 
    "crypto"
    "strings"
    "github.com/astaxie/beego/logs"
    "github.com/byyh/go/com"
    )

type AlipayBase struct {
    AppId string    
    RsaPrivateKey string
    AlipayRsaPublicKey  string
    ApiVersion  string   // '1.0'
    Charset  string  //  'UTF-8'
    FileCharset  string  //  'UTF-8'
    Format   string      // 'json'
    SignType  string
    Sign   string   // 
    BizContent  string
    Url  string
    Client  * http.Client
    Method   string
    Timestamp string   // Y-m-d H:i:s
    PostUrl  string
    GatewayUrl  string
}

func (this *AlipayBase) Post() string {   
    strContent := this.CompareSignParam()
    this.Sign = this.getSign(strContent)

    urlStr := this.GatewayUrl + this.UrlParse()

    fmt.Println(urlStr)
    bodyStr := strings.NewReader("biz_content=" + this.BizContent)
    
    this.Client = &http.Client{}
  
    req, _ := http.NewRequest("POST", urlStr, bodyStr)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

    resp, err := this.Client.Do(req)  // 发送
    defer resp.Body.Close()      // 关闭resp.Body
    if err != nil { 
        fmt.Println("发送失败", err.Error()) 

        return ""
    }    

    if 200 == resp.StatusCode {
        body, _ := ioutil.ReadAll(resp.Body) 

        return string(body)
    } else {
        fmt.Println("network error", resp.StatusCode)
    }

    return ""
}

func (this *AlipayBase) CompareSignParam() (string) {
    err := this.CheckParam()
    if nil != err {
        panic(err)
    }

    str := "app_id=" + this.AppId
    //str += "&biz_content=" + url.QueryEscape(this.BizContent)
    str += "&biz_content=" + this.BizContent
    str += "&charset=" + this.Charset
    str += "&format=" + this.Format
    str += "&method=" + this.Method
    str += "&sign_type=" + this.SignType
    //str += "&timestamp=" + url.QueryEscape(this.Timestamp)
    str += "&timestamp=" + this.Timestamp
    str += "&version=" + this.ApiVersion

    return  str; 
}

func (this *AlipayBase) UrlParse() (strUrlParam string) {
    strUrlParam = "app_id=" + this.AppId
    strUrlParam += "&charset=" + this.Charset
    strUrlParam += "&format=" + this.Format
    strUrlParam += "&method=" + this.Method
    strUrlParam += "&sign_type=" + this.SignType
    strUrlParam += "&timestamp=" + url.QueryEscape(this.Timestamp)
    strUrlParam += "&version=" + this.ApiVersion
    strUrlParam += "&sign=" + url.QueryEscape(this.Sign)

    return
}

func (this *AlipayBase) ToXml() string {
    err := this.CheckParam()
    if nil != err {
        logs.Error("ToXml failed", err)
        panic(err)
    }

    xml := "<xml>"
    xml += "<appid>" + this.AppId + "</appid>"
    xml += "<sign>" + this.getSign(this.AppId) + "</sign>"
    xml += "</xml>"

    return  xml; 
}

func (this *AlipayBase) CheckParam() (error) {
    if("" == this.AppId) {
        return errors.New("AppId is not allow empty") 
    }

    if("" == this.RsaPrivateKey) {
        return errors.New("RsaPrivateKey is not allow empty") 
    }

    if("" == this.AlipayRsaPublicKey) {
        return errors.New("AlipayrsaPublicKey is not allow empty") 
    }

    if("" == this.ApiVersion) {
        this.ApiVersion = "1.0"
    }

    if("" == this.Charset) {
        this.Charset = "UTF-8"
    }

    if("" == this.Format) {
        this.Format = "json"
    }

    if("" == this.SignType) {
        return errors.New("SignType is not allow empty") 
    }

    if("" == this.Timestamp) {
        return errors.New("Timestamp is not allow empty") 
    }

    return nil
}

func (this *AlipayBase) getSign(str string) (string) {  
    if !strings.Contains(this.RsaPrivateKey, "BEGIN PRIVATE KEY") {
        this.RsaPrivateKey = 
            "-----BEGIN PRIVATE KEY-----\n" +
            this.RsaPrivateKey + 
            "\n-----END PRIVATE KEY-----"
    }

    strRsa2 := com.SignRsa([]byte(str), this.RsaPrivateKey, crypto.SHA256)

    return strRsa2;
}
