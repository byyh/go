package alipay 

import ( 
    "fmt" 
    "encoding/json"
    )

type PayBill struct {
    AlipayBase 
    RequestParam  BillMemberVar     
}

type BillMemberVar struct {        
    BillDate  string   `json:"bill_date,omitempty"`
    BillType  string   `json:"bill_type,omitempty"`
}

func (this *PayBill) Post() (string) {
    // 
    this.Method = "alipay.data.dataservice.bill.downloadurl.query"
    this.GatewayUrl = "https://openapi.alipay.com/gateway.do?"

    this.setRequestParam();

    result := this.AlipayBase.Post()

    return result;
}

func (this *PayBill) setRequestParam()  {
    // 
    data, err := json.Marshal(this.RequestParam)
    if nil != err {
        fmt.Println("设置参数错误:", err)
        panic(err)
    }

    this.BizContent = string(data)
}