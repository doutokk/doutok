package alipay

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/doutokk/doutok/app/payment/conf"
	"github.com/smartwalle/alipay/v3"
	"net/url"
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
	p.NotifyURL = c.CallBackUrl
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

type CallbackParams struct {
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

func VerifyCallback(params CallbackParams) bool {
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
