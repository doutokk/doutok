package alipay

import (
	"context"
	"testing"
)

func Test_createPayOrder(t *testing.T) {
	type args struct {
		orderId string
		amount  float64
	}
	tests := []struct {
		name    string
		args    args
		wantUrl string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				orderId: "2014112611001004680073956707",
				amount:  1,
			},
			wantUrl: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUrl, err := CreatePayOrder(tt.args.orderId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePayOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUrl != tt.wantUrl {
				t.Errorf("CreatePayOrder() gotUrl = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}

func Test_verifyCallback(t *testing.T) {
	type args struct {
		params CallbackParams
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				params: CallbackParams{
					Charset:     "utf-8",
					OutTradeNo:  "2025-02-27 13:08:47.1056417 +0800 CST m=+0.005788301",
					Method:      "alipay.trade.page.pay.return",
					TotalAmount: "1.00",
					Sign:        "EbNxV0pj0luLDTRYbpPsoxjg+Rb4oDjja1FA9jQO/z4MRPSbKrj7Ts6qOPi+O2T2j6+SgnuoGE2UjJV7+mKgr/z14VwFB0O8QnerPqBUU414RKNX3vM9SRXNbOSEUWc2N2Nuux96hBPxOVvWigTylLyIw58iqBlFus3+ImIwZC8rte1Ybf/67iN6ay9oecYvtjDkJLgBm/Yw6FDc5nBpakIGmO6BKHyTu1K5qz58LI4PA/OweKm2By9Cm2scqZDPGmCy3TDzuK2xc51O5heFlb5Ig/RhN8efIaVuokXYgqD9/P03sS1c7kvB+44KePt9RpsRhtZtGWGkICHKgHVOLw==",
					TradeNo:     "2025022722001463260505380232",
					AuthAppId:   "9021000134679168",
					Version:     "1.0",
					AppId:       "9021000134679168",
					SignType:    "RSA2",
					SellerId:    "2088721030798344",
					Timestamp:   "2025-02-27 13:09:55",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyCallback(tt.args.params); got != tt.want {
				t.Errorf("VerifyCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyNotify(t *testing.T) {
	type args struct {
		ctx       context.Context
		partnerId string
		notifyId  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				ctx:       context.TODO(),
				partnerId: "",
				notifyId:  "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyNotify(tt.args.ctx, tt.args.partnerId, tt.args.notifyId); got != tt.want {
				t.Errorf("verifyNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTradeQuery(t *testing.T) {
	type args struct {
		ctx     context.Context
		tradeNo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				tradeNo: "2025022722001463260505368952",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := TradeQuery(tt.args.ctx, tt.args.tradeNo); got != tt.want {
				t.Errorf("TradeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
