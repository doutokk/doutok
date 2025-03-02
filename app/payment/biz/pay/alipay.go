package pay

import (
	"context"
	_ "embed"
	"fmt"
	"net/url"

	"github.com/doutokk/doutok/app/payment/conf"
	"github.com/smartwalle/alipay/v3"
)

var (
	//go:embed alipayPublicCert.crt
	alipayPublicCert []byte
	//go:embed alipayRootCert.crt
	alipayRootCert []byte
	//go:embed appPublicCert.crt
	appPublicCert []byte

	c         = conf.GetConf().Alipay
	client, _ = alipay.New(c.AppID, c.PrivateKey, false)
)

func init() {
	var err error
	err = client.LoadAppCertPublicKey(string(appPublicCert))
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayRootCert(string(alipayRootCert))
	if err != nil {
		panic(err)
	}
	err = client.LoadAlipayCertPublicKey(string(alipayPublicCert))
	if err != nil {
		panic(err)
	}
}

func CreatePayOrder(orderId string, amount float64) (url string, err error) {
	p := alipay.TradePagePay{}
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.NotifyURL = c.NotifyBackUrl
	p.ReturnURL = c.CallBackUrl
	p.Subject = "Order: " + orderId
	p.OutTradeNo = orderId
	p.TotalAmount = fmt.Sprintf("%.2f", amount)

	result, err := client.TradePagePay(p)
	if err != nil {
		return
	}
	url = result.String()

	return
}

type ReturnCallbackParams struct {
	Charset     string
	OutTradeNo  string
	Method      string
	TotalAmount string
	Sign        string
	TradeNo     string
	AuthAppId   string
	Version     string
	AppId       string
	SignType    string
	SellerId    string
	Timestamp   string
}

// AlipayCallbackNotification represents the callback notification from Alipay payment system
type AlipayCallbackNotification struct {
	GmtCreate      string `form:"gmt_create" json:"gmt_create"`             // Creation time
	Charset        string `form:"charset" json:"charset"`                   // Character encoding
	GmtPayment     string `form:"gmt_payment" json:"gmt_payment"`           // Payment time
	NotifyTime     string `form:"notify_time" json:"notify_time"`           // Notification time
	Subject        string `form:"subject" json:"subject"`                   // Order subject/title
	Sign           string `form:"sign" json:"sign"`                         // Signature
	BuyerId        string `form:"buyer_id" json:"buyer_id"`                 // Buyer ID
	InvoiceAmount  string `form:"invoice_amount" json:"invoice_amount"`     // Invoice amount
	Version        string `form:"version" json:"version"`                   // API version
	NotifyId       string `form:"notify_id" json:"notify_id"`               // Notification ID
	FundBillList   string `form:"fund_bill_list" json:"fund_bill_list"`     // Payment source details (JSON string)
	NotifyType     string `form:"notify_type" json:"notify_type"`           // Notification type
	OutTradeNo     string `form:"out_trade_no" json:"out_trade_no"`         // Merchant order number
	TotalAmount    string `form:"total_amount" json:"total_amount"`         // Total amount
	TradeStatus    string `form:"trade_status" json:"trade_status"`         // Trade status
	TradeNo        string `form:"trade_no" json:"trade_no"`                 // Alipay trade number
	AuthAppId      string `form:"auth_app_id" json:"auth_app_id"`           // Authorized app ID
	ReceiptAmount  string `form:"receipt_amount" json:"receipt_amount"`     // Receipt amount
	PointAmount    string `form:"point_amount" json:"point_amount"`         // Point amount used
	BuyerPayAmount string `form:"buyer_pay_amount" json:"buyer_pay_amount"` // Buyer payment amount
	AppId          string `form:"app_id" json:"app_id"`                     // App ID
	SignType       string `form:"sign_type" json:"sign_type"`               // Signature type
	SellerId       string `form:"seller_id" json:"seller_id"`               // Seller ID
}

func (n *AlipayCallbackNotification) GetSignData() url.Values {
	var v = url.Values{}
	v.Add("gmt_create", n.GmtCreate)
	v.Add("charset", n.Charset)
	v.Add("gmt_payment", n.GmtPayment)
	v.Add("notify_time", n.NotifyTime)
	v.Add("subject", n.Subject)
	v.Add("sign", n.Sign)
	v.Add("buyer_id", n.BuyerId)
	v.Add("invoice_amount", n.InvoiceAmount)
	v.Add("version", n.Version)
	v.Add("notify_id", n.NotifyId)
	v.Add("fund_bill_list", n.FundBillList)
	v.Add("notify_type", n.NotifyType)
	v.Add("out_trade_no", n.OutTradeNo)
	v.Add("total_amount", n.TotalAmount)
	v.Add("trade_status", n.TradeStatus)
	v.Add("trade_no", n.TradeNo)
	v.Add("auth_app_id", n.AuthAppId)
	v.Add("receipt_amount", n.ReceiptAmount)
	v.Add("point_amount", n.PointAmount)
	v.Add("buyer_pay_amount", n.BuyerPayAmount)
	v.Add("app_id", n.AppId)
	v.Add("sign_type", n.SignType)
	v.Add("seller_id", n.SellerId)
	return v
}

func VerifyReturnCallback(params ReturnCallbackParams) bool {
	var v = url.Values{}

	// Convert struct fields to url.Values
	v.Set("charset", params.Charset)
	v.Set("out_trade_no", params.OutTradeNo)
	v.Set("method", params.Method)
	v.Set("total_amount", params.TotalAmount)
	v.Set("sign", params.Sign)
	v.Set("trade_no", params.TradeNo)
	v.Set("auth_app_id", params.AuthAppId)
	v.Set("version", params.Version)
	v.Set("app_id", params.AppId)
	v.Set("sign_type", params.SignType)
	v.Set("seller_id", params.SellerId)
	v.Set("timestamp", params.Timestamp)

	err := client.VerifySign(v)
	if err == nil {
		return true
	}
	return false
}

func VerifyNotifyCallback(ctx context.Context, n AlipayCallbackNotification) bool {
	// DecodeNotification 内部已调用 VerifySign 方法验证签名
	var _, err = client.DecodeNotification(n.GetSignData())
	if err != nil {
		// 错误处理
		fmt.Println(err)
		return false
	}
	return true
}

func verifyNotify(ctx context.Context, partnerId string, notifyId string) bool {

	return client.NotifyVerify(ctx, partnerId, notifyId)
}

func TradeQuery(ctx context.Context, tradeNo string) (status string, ok bool) {

	param := alipay.TradeQuery{
		TradeNo: tradeNo,
	}
	result, err := client.TradeQuery(ctx, param)
	if err != nil {
		return "", false
	}
	status = string(result.TradeStatus)
	return status, status == "TRADE_SUCCESS"
}

func CancelOrder(tradeNo string) (ok bool) {
	param := alipay.TradeCancel{
		OutTradeNo: tradeNo,
	}
	_, err := client.TradeCancel(context.Background(), param)
	if err != nil {
		return false
	}
	return true
}
