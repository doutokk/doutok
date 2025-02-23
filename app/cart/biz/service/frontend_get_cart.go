package service

import (
	"context"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

type FrontendGetCartService struct {
	ctx context.Context
}

// NewFrontendGetCartService new FrontendGetCartService
func NewFrontendGetCartService(ctx context.Context) *FrontendGetCartService {
	return &FrontendGetCartService{ctx: ctx}
}

// Run create note info
func (s *FrontendGetCartService) Run(req *cart.FrontendGetCartReq) (resp *cart.FrontendGetCartResp, err error) {
	// 不确定grpc gateway转发的过程中ctx会不会丢失。。。
	//userId := utils.GetUserIdFromCtx(s.ctx)
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//server := service.NewGetCartService(ctx)
	//resp, err := server.Run(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//// 下面是bff服务提供
	//productCli := rpc.ProductClient
	//var bffCart []*BffProduct
	//items := resp.Cart.Items
	//for _, item := range items {
	//	getProductReq := &product.GetProductReq{Id: item.ProductId}
	//	productResp, productErr := productCli.GetProduct(ctx, getProductReq)
	//	if productErr != nil {
	//		c.String(consts.StatusBadRequest, productErr.Error())
	//		return
	//	}
	//	bffProduct := productResp.Product
	//	bffCart = append(bffCart, &BffProduct{
	//		Description: bffProduct.Description,
	//		Img:         bffProduct.Picture,
	//		Price:       bffProduct.Price,
	//		ProductID:   bffProduct.Id,
	//		ProductName: bffProduct.Name,
	//		Quantity:    item.Quantity,
	//	})
	//}
	//
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//// cart
	//
	////
	//
	//c.JSON(consts.StatusOK, bffCart)

	return
}
