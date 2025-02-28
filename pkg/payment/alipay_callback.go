package payment

import "net/url"

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
