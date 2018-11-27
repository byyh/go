package thirdPay 

import ( 
    "bytes" 
    "errors"
    "fmt" 
    "io/ioutil" 
    "net/http" 
    //"net/url" 
    //"strings"
    // "os" 
    "github.com/byyh/go/com"
    ) 

type WeixinPay struct {
    Key string
    Appid string
    BillDate string
    BillType string
    MchId string   
    NonceStr string
    Sign   string
}

var (
    
)

func (this *WeixinPay) PostDateBill() string {
    const BILL_POST_URL string = "https://api.mch.weixin.qq.com/pay/downloadbill"

    client := &http.Client{}

    xml := this.ToXml()

    body := bytes.NewReader([]byte(xml))    
    req, _ := http.NewRequest("POST", BILL_POST_URL, body)

    req.Header.Set("Content-Type", "text/xml") 
    //fmt.Println(xml, req)

    resp, err := client.Do(req)  // 发送
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

func (this *WeixinPay) ToXml() string {
    err := this.CheckBillParam()
    if nil != err {
        panic(err)
    }

    xml := "<xml>"
    xml += "<appid>" + this.Appid + "</appid>"
    xml += "<bill_date>" + this.BillDate + "</bill_date>"
    xml += "<bill_type>" + this.BillType + "</bill_type>"
    xml += "<mch_id>" + this.MchId + "</mch_id>"    
    xml += "<nonce_str>" + this.NonceStr + "</nonce_str>"
    xml += "<sign>" + this.ToSign() + "</sign>"
    xml += "</xml>"

    return  xml; 
}

func (this *WeixinPay) CheckBillParam() (error) {
    if("" == this.Appid) {
        return errors.New("Appid is not allow empty") 
    }

    if("" == this.MchId) {
        return errors.New("MchId is not allow empty") 
    }

    if("" == this.BillType) {
        return errors.New("BillType is not allow empty") 
    }

    if("" == this.BillDate) {
        return errors.New("BillDate is not allow empty") 
    }

    if("" == this.NonceStr) {
        return errors.New("NonceStr is not allow empty") 
    }

    return nil
}

func (this *WeixinPay) ToSign() (string) {

    str := "appid=" + this.Appid + "&bill_date=" + this.BillDate + "&bill_type=" + this.BillType + 
        "&mch_id=" + this.MchId + 
        "&nonce_str=" + this.NonceStr +
        "&key=" + this.Key

    strMd5 := com.Md5(str)    
    //return com.Sha256Code(strMd5, this.Key)
    return strMd5;
}
