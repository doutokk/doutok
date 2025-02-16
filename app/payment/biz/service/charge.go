package service

import (
	"context"
	"github.com/doutokk/doutok/app/payment/biz/dal/model"
	"github.com/doutokk/doutok/app/payment/biz/dal/mysql"
	payment "github.com/doutokk/doutok/app/payment/kitex_gen/payment"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	_ = godotenv.Load()
	//todo: 略检测信用卡有效，扣款
	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: translationId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
}
